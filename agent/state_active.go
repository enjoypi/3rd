package agent

import (
	"errors"
	"io"
	"net"

	sc "github.com/enjoypi/gostatechart"
)

type StateActive struct {
	sc.SimpleState
}

func (state *StateActive) Begin(context interface{}, event sc.Event) sc.Event {
	return nil
}

func (state *StateActive) End(event sc.Event) sc.Event {
	return nil
}

func (state *StateActive) GetTransitions() sc.Transitions {
	trans := sc.NewTranstions()
	trans.RegisterTransition(io.EOF, (*stateClosing)(nil))
	trans.RegisterTransition((*net.OpError)(nil), (*stateClosing)(nil))
	trans.RegisterTransition(errors.New(""), (*stateClosing)(nil))
	//trans.RegisterTransition((*evKickByGame)(nil), (*stateClosing)(nil))
	return trans
}

func (state *StateActive) InitialChildState() sc.State {
	return (*stateAuth)(nil)
}
