package autonomous

import (
	"sync"
)

type AgentHub struct {
	Life
	Stopper
	Managed

	start            chan Agent
	stop             chan Agent
	registeredAgents map[Agent]bool
}

func NewAgentHub() *AgentHub {
	h := new(AgentHub)
	h.Life = NewLife()
	return h
}

func (h *AgentHub) StartAgent(a Agent) {
	h.start <- a
}

func (h *AgentHub) StopAgent(a Agent) {
	h.stop <- a
}

func (h *AgentHub) Run() {
	h.startup()
	h.Life.Begin()

Run:
	for {
		select {
		case a := <-h.start:
			go a.Run()
			h.registeredAgents[a] = true
		case a := <-h.stop:
			go a.Stop()
			delete(h.registeredAgents, a)
		case <-h.Stopper:
			break Run
		}
	}

	h.shutdown()
	h.Life.End()
}

func (h *AgentHub) startup() {
}

func (h *AgentHub) shutdown() {
	var wg sync.WaitGroup

	for a, _ := range h.registeredAgents {
		wg.Add(1)
		go a.Stop()
		go func() {
			a.Stopped().Wait()
			wg.Done()
		}()
	}

	wg.Wait()
}
