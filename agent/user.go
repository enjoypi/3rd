package agent

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
)

type UserAgent struct {
	actors.SimpleActor
}

func init() {
	actors.RegisterActorCreator("agent", NewUserAgent)
}

func NewUserAgent() actors.Actor {
	return &UserAgent{}
}

func (u *UserAgent) Handle(message types.Message) error {
	return nil
}

func (u *UserAgent) Initialize(v *viper.Viper) error {
	_ = u.SimpleActor.Initialize()
	return nil
}

func (u *UserAgent) Terminate() {

}
