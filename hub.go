package autonomous

import (
	"sync"
)

type Hub struct {
	Life
	Stopper
	Managed

	start  chan Agent
	stop   chan Agent
	agents map[Agent]bool
}

func NewHub() *Hub {
	h := new(Hub)
	h.Life = NewLife()
	return h
}

func (h *Hub) StartAgent(a Agent) {
	h.start <- a
}

func (h *Hub) StopAgent(a Agent) {
	h.stop <- a
}

func (h *Hub) Run() {
	h.Life.Begin()

Run:
	for {
		select {
		case a := <-h.start:
			go a.Start()
			h.agents[a] = true
		case a := <-h.stop:
			go a.Stop()
			delete(h.agents, a)
		case <-h.Stopper:
			break Run
		}
	}

	h.shutdown()
	h.Life.End()
}

func (h *Hub) shutdown() {
	var wg sync.WaitGroup

	for agent, _ := range h.agents {
		wg.Add(1)
		go agent.Stop()
		go func() {
			agent.Stopped().Wait()
			wg.Done()
		}()
	}

	wg.Wait()
}

func (h *Hub) Agents() map[Agent]bool {
	as := h.agents
	return as
}
