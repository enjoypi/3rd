package agent

import (
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"

	"github.com/enjoypi/god/stdlib"
)

type UserAgent struct {
	stdlib.DefaultActor
}

func init() {
	stdlib.RegisterActorCreator("agent", NewUserAgent)
}

func NewUserAgent() stdlib.Actor {
	return &UserAgent{}
}

func (u *UserAgent) Handle(message types.Message) error {
	return nil
}

func (u *UserAgent) Initialize(v *viper.Viper) error {
	_ = u.DefaultActor.Initialize()
	return nil
}

func (u *UserAgent) Terminate() {

}
