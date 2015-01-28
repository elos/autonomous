package autonomous

import (
	"sync"
)

type NullHub struct {
	*Core
	m                sync.Mutex
	RegisteredAgents map[Agent]bool
}

func NewNullHub() *NullHub {
	return &NullHub{
		Core:             NewCore(),
		RegisteredAgents: make(map[Agent]bool),
	}
}

func (h *NullHub) StartAgent(a Agent) {
	h.m.Lock()
	defer h.m.Unlock()

	h.RegisteredAgents[a] = true
}

func (h *NullHub) StopAgent(a Agent) {
	h.m.Lock()
	defer h.m.Unlock()

	delete(h.RegisteredAgents, a)
}

func (h *NullHub) Reset() {
	h.m.Lock()
	defer h.m.Unlock()

	h.Core.Shutdown()
	h.RegisteredAgents = make(map[Agent]bool)
}
