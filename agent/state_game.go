package agent

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/enjoypi/god"

	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type stateGame struct {
	sc.SimpleState

	//*net.Session
	agentSub *nats.Subscription
}

func (state *stateGame) Begin(context interface{}, event sc.Event) sc.Event {
	//state.Session = context.(*net.Session)

	state.registerReactions()

	var err error
	//state.agentSub, err = state.Node.Subscribe("agent.1.*", state.onNatsMsg)
	//state.agentSub, err = state.Node.Subscribe(">", state.onNatsMsg)
	if err != nil {
		return err
	}

	return nil
}

func (state *stateGame) GetEvent() sc.Event {
	//session := state.Session

	var header pb.Header
	//if err := session.RecvMsg(&header); err != nil {
	//	return err
	//}

	//typ, ok := p.id2type[header.MessageType]
	//if !ok {
	//	return 0, nil, ErrMessageNotRegistered
	//}
	//// 根据类型创建一个对应的实例
	//msg0 := reflect.New(typ.Elem()).Interface().(proto.Message)

	var req proto.Message
	switch header.MessageType {
	case "pb.Heartbeat":
		//msg0 := reflect.New(reflect.TypeOf(pb.Heartbeat{}).Elem()).Interface().(proto.Message)
		req = &pb.Heartbeat{}
	case "pb.Echo":
		req = &pb.Echo{}
	default:
		return nil
	}

	//if err := session.RecvMsg(req); err != nil {
	//	return err
	//}

	//session.Info(req.String(), zap.String("type", reflect.TypeOf(req).String()))
	if !state.HasReaction(req) {
		// TODO Post To Manager
		return nil
	}
	return req
}

func (state *stateGame) registerReactions() {
	state.RegisterReaction((*pb.Echo)(nil), state.onEcho)
	state.RegisterReaction((*pb.Heartbeat)(nil), state.onHeartbeat)
}

func (state *stateGame) onEcho(event sc.Event) sc.Event {
	//req := event.(*pb.Echo)
	//return state.Node.CastTo(pb.ServiceType_Mesh, req)
	//return state.Session.SendMsg(req)
	return nil
}

func (state *stateGame) onHeartbeat(event sc.Event) sc.Event {
	req := event.(*pb.Heartbeat)
	req.ToTimestamp = time.Now().UnixNano()
	//return state.Session.SendMsg(req)
	return nil
}

func (state *stateGame) onNatsMsg(msg *nats.Msg) {
	god.Logger.Debug("onNatsMsg", zap.String("subject", msg.Subject), zap.Int("size", len(msg.Data)))
	buf := bytes.NewBuffer(msg.Data)
	var l uint16
	if err := binary.Read(buf, binary.LittleEndian, &l); err != nil {
		god.Logger.Warn("onNatsMsg error", zap.Error(err))
		return
	}
	sizeofLen := 2

	var header pb.Header
	if err := header.Unmarshal(msg.Data[sizeofLen : sizeofLen+int(l)]); err != nil {
		god.Logger.Warn("NATS error", zap.Error(err))
		return
	}
	god.Logger.Debug("header", zap.String("type", header.MessageType))

	switch header.MessageType {
	case "pb.Heartbeat":
		//msg0 := reflect.New(reflect.TypeOf(pb.Heartbeat{}).Elem()).Interface().(proto.Message)
		req := &pb.Heartbeat{}
		if err := req.Unmarshal(msg.Data[sizeofLen+int(l)+sizeofLen:]); err != nil {
			god.Logger.Warn("NATS Unmarshal failed", zap.Error(err), zap.String("msgType", header.MessageType))
			return
		}
		state.Outermost().PostEvent(req)
	case "pb.Echo":
		req := &pb.Echo{}
		if err := req.Unmarshal(msg.Data[sizeofLen+int(l)+sizeofLen:]); err != nil {
			god.Logger.Warn("NATS Unmarshal failed", zap.Error(err), zap.String("msgType", header.MessageType))
			return
		}
		state.Outermost().PostEvent(req)
	default:
		return
	}

}
