package agent

import (
	"github.com/enjoypi/god/services/net"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

type stateClosing struct {
	sc.SimpleState
}

func (s *stateClosing) Begin(context interface{}, event sc.Event) sc.Event {
	session := context.(*net.Session)
	session.Logger.Info("closing", zap.Any("event", event))
	//<-session.Context().Done()
	s.Outermost().Terminate(event)
	return nil
}
