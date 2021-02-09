package agent

import (
	"math/rand"

	"github.com/enjoypi/god/core"
)

type UserAgent struct {
	core.DefaultActor
}

var (
	userAgentType = rand.Int63()
)

func init() {
	core.RegisterActorCreator(userAgentType, NewUserAgent)
}

func NewUserAgent() core.Actor {
	return &UserAgent{}
}

func (u *UserAgent) Handle(message core.Message) core.Message {
	return nil
}

func (u *UserAgent) Initialize() error {
	_ = u.DefaultActor.Initialize()
	return nil
}

func (u *UserAgent) Terminate() {

}
