package autonomous

import (
	"sync"
)

type NullHub struct {
	Life
	Stopper
	Managed

	m                sync.Mutex
	RegisteredAgents map[Agent]bool
}

func NewNullHub() *NullHub {
	h := new(NullHub)
	h.Life = NewLife()
	return h
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

	h.Stop()
	h.RegisteredAgents = make(map[Agent]bool)
}
