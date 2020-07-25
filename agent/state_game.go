package agent

import (
	"reflect"
	"time"

	"github.com/enjoypi/god/pb"
	"github.com/enjoypi/god/services/net"
	sc "github.com/enjoypi/gostatechart"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

type stateGame struct {
	sc.SimpleState

	*net.Session
}

func (state *stateGame) Begin(context interface{}, event sc.Event) sc.Event {
	state.Session = context.(*net.Session)

	state.registerReactions()

	return nil
}

func (state *stateGame) GetEvent() sc.Event {
	session := state.Session

	var header pb.Header
	if err := session.RecvMsg(&header); err != nil {
		return err
	}

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

	if err := session.RecvMsg(req); err != nil {
		return err
	}

	session.Info(req.String(), zap.String("type", reflect.TypeOf(req).String()))
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
	req := event.(*pb.Echo)
	return state.Session.SendMsg(req)
}

func (state *stateGame) onHeartbeat(event sc.Event) sc.Event {
	req := event.(*pb.Heartbeat)
	req.ToTimestamp = time.Now().UnixNano()
	return state.Session.SendMsg(req)
}
