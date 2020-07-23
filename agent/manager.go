package agent

import (
	"github.com/enjoypi/god/pb"
	"github.com/enjoypi/god/services/net"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

type Manager struct {
	sc.SimpleState
	*net.Service
}

func (m *Manager) Begin(context interface{}, event sc.Event) sc.Event {
	m.Service = context.(*net.Service)
	m.RegisterReaction((*pb.Header)(nil), m.onHeader)
	return nil
}

func (m *Manager) GetTransitions() sc.Transitions {
	return nil
}

func (m *Manager) onHeader(event sc.Event) sc.Event {
	m.Logger.Info("onHeader", zap.Any("event", event))
	return nil
}
