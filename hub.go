package autonomous

import "sync"

type Hub struct {
	Life
	Stopper
	Managed

	start   chan Agent
	stop    chan Agent
	agents  map[Agent]bool
	mapLock *sync.Mutex
}

func NewHub() *Hub {
	h := &Hub{
		Life:    NewLife(),
		Stopper: make(Stopper),
		Managed: *new(Managed),
		start:   make(chan Agent),
		stop:    make(chan Agent),
		agents:  make(map[Agent]bool),
		mapLock: new(sync.Mutex),
	}

	return h
}

func (h *Hub) StartAgent(a Agent) {
	h.start <- a
}

func (h *Hub) StopAgent(a Agent) {
	h.stop <- a
}

func (h *Hub) Start() {
	h.Life.Begin()

Run:
	for {
		select {
		case a := <-h.start:
			go a.Start()
			h.mapLock.Lock()
			h.agents[a] = true
			h.mapLock.Unlock()
		case a := <-h.stop:
			go a.Stop()
			h.mapLock.Lock()
			delete(h.agents, a)
			h.mapLock.Unlock()
		case <-h.Stopper:
			break Run
		}
	}

	h.shutdown()
}

func (h *Hub) shutdown() {
	var wg sync.WaitGroup

	h.mapLock.Lock()
	for agent, _ := range h.agents {
		wg.Add(1)
		go func() {
			wg.Done()
			agent.WaitStop()
			wg.Done()
		}()
	}

	wg.Wait()

	for agent, _ := range h.agents {
		wg.Add(1) // for each post-wait stop
		go func() {
			agent.Stop()
		}()
	}
	h.mapLock.Unlock()

	wg.Wait()
	h.Life.End()
}

func (h *Hub) Agents() map[Agent]bool {
	as := make(map[Agent]bool)
	h.mapLock.Lock()
	for k, v := range h.agents {
		as[k] = v
	}
	h.mapLock.Unlock()
	return as
}
