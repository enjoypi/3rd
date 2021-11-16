package agent

import (
	"context"

	"github.com/enjoypi/god/logger"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

type stateClosing struct {
	sc.SimpleState
}

func (s *stateClosing) Begin(ctx context.Context, event sc.Event) sc.Event {
	//session := context.(*net.Session)
	logger.L.Info("closing", zap.Any("event", event))
	//<-session.Context().Done()
	s.Outermost().Terminate(event)
	return nil
}
