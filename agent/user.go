package agent

import (
	"github.com/enjoypi/god/def"
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

func (u *UserAgent) Handle(message def.Message) error {
	return nil
}

func (u *UserAgent) Initialize(v *viper.Viper) error {
	_ = u.SimpleActor.Initialize()
	return nil
}

func (u *UserAgent) Terminate() {

}
