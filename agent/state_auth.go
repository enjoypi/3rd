package agent

import (
	"context"

	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
)

type stateAuth struct {
	sc.SimpleState

	//*net.Session
}

func (state *stateAuth) GetTransitions() sc.Transitions {
	trans := sc.NewTranstions()
	trans.RegisterTransition((*pb.AuthAck)(nil), (*stateGame)(nil))
	return trans
}

func (state *stateAuth) Begin(ctx context.Context, event sc.Event) sc.Event {
	//state.Session = context.(*net.Session)
	//var header pb.Header
	//if err := state.Session.RecvMsg(&header); err != nil {
	//	return err
	//}
	var req pb.AuthReq
	//if err := state.Session.RecvMsg(&req); err != nil {
	//	return err
	//}
	return state.onAuthReq(ctx, &req)
}

func (state *stateAuth) onAuthReq(ctx context.Context, req *pb.AuthReq) sc.Event {
	//session := state.Session

	var ack pb.AuthAck
	//defer session.SendMsg(&ack)

	// signature := req.Signature
	// TODO 到Entrance验签
	//if req.CInfo == nil {
	//	ack.ErrCode = pb.ErrCodeNilClientInfo
	//	return errors.New("invalid client info")
	//}

	// 发送到平台进行权证验证
	//ri := god.Call(conf.PlatformServcie, 0,
	//	&pfproto.PfAuthReq{
	//		AuthType:   req.UserType,
	//		User:       req.UserName,
	//		Password:   req.Password,
	//		DeviceID:   req.CInfo.OpenUDID,
	//		IsRegister: false,
	//	})
	//if ri.Err != nil {
	//	ack.ErrCode = pb.ErrCodeFramework
	//	return ri.Err
	//}

	// 验证结果封装
	//pfAck := ri.Ack.(*pfproto.PfAuthAck)
	//if pfAck.PfErr != nil {
	//	ack.ErrCode = pb.ErrCodePfAuthFailed
	//	return pfAck.PfErr
	//}

	// 封装角色列表
	//characters := make([]*pb.Character, len(pfAck.Characters))
	//for i, character := range pfAck.Characters {
	//	// 解析meta信息
	//	meta := &pb.CharacterMeta{}
	//}

	//// 获取服务器列表
	//servers, code := agent.packServers()
	//if code != pb.ErrCodeSuccess {
	//	agent.SendToClient(&pb.AuthAck{ErrCode: code})
	//}
	// 返回
	//authAck := &pb.AuthAck{
	//	ErrCode:     pb.ErrCodeSuccess,
	//	AccessToken: pfAck.AccessToken,
	//	AccountID:   pfAck.AccountID,
	//	Type:        req.UserType,
	//	Characters:  characters,
	//	ClientState: req.ClientState,
	//	Servers:     servers,
	//	RecommandServer: selectRecommand(servers),
	//	ExpireTs: pfAck.ExpireAt,
	//}

	return &ack
}
