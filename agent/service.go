package agent

import "github.com/enjoypi/god/core"

type Service struct {
}

func (s Service) Name() string {
	return "agent"
}

func (s Service) Start() error {
	return nil
}

func (s Service) Stop() {
}

func NewService() core.Service {
	return &Service{}
}
