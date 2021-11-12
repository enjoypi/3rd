package agent

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/def"
	"github.com/spf13/viper"
)

type UserAgent struct {
	actors.SimpleActor
}

const atAgent = def.ATUser + 1

func init() {
	actors.RegisterActorCreator(atAgent, NewUserAgent)
}

func NewUserAgent() actors.Actor {
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
