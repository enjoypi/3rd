package agent

import (
	"context"
	"net"
	"time"

	"github.com/enjoypi/god/logger"
	"go.uber.org/zap"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type UserAgent struct {
	stdlib.SimpleActor
}

const atAgent = "agent"

func init() {
	stdlib.RegisterActorCreator(atAgent, NewUserAgent)
}

func NewUserAgent() stdlib.Actor {
	return &UserAgent{}
}

func (u *UserAgent) Initialize(v *viper.Viper) error {
	_ = u.SimpleActor.Initialize()
	u.RegisterReaction((*event.EvSocketConnected)(nil), u.onSocketConnected)
	return nil
}

func (u *UserAgent) onSocketConnected(ctx context.Context, message def.Message, args ...interface{}) def.Message {
	conn := message.(*event.EvSocketConnected).Conn
	logger.L.Debug("new user",
		//zap.String("network", conn.LocalAddr().Network()),
		//zap.String("local", conn.LocalAddr().String()),
		zap.String("remote", conn.RemoteAddr().String()),
	)

	tcp := conn.(*net.TCPConn)
	tcp.CloseWrite()
	time.Sleep(10 * time.Second)
	tcp.CloseRead()
	return nil
}

func (u *UserAgent) Terminate() {

}
