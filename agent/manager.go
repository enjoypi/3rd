package agent

import (
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

type Manager struct {
	sc.SimpleState
}

func (m *Manager) Begin(context interface{}, event sc.Event) sc.Event {
	m.RegisterReaction((*pb.Header)(nil), m.onHeader)
	return nil
}

func (m *Manager) onHeader(event sc.Event) sc.Event {
	logger.L.Info("onHeader", zap.Any("event", event))
	return nil
}
