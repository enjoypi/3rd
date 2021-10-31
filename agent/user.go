package agent

import (
	"math/rand"

	"github.com/enjoypi/god/types"

	"github.com/enjoypi/god/stdlib"
)

type UserAgent struct {
	stdlib.DefaultActor
}

var (
	userAgentType = rand.Int63()
)

func init() {
	stdlib.RegisterActorCreator(userAgentType, NewUserAgent)
}

func NewUserAgent() stdlib.Actor {
	return &UserAgent{}
}

func (u *UserAgent) Handle(message types.Message) types.Message {
	return nil
}

func (u *UserAgent) Initialize() error {
	_ = u.DefaultActor.Initialize()
	return nil
}

func (u *UserAgent) Terminate() {

}
