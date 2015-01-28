package autonomous

type AgentHub struct {
	*Core

	start            chan Agent
	stop             chan Agent
	registeredAgents map[Agent]bool
}

func NewAgentHub() *AgentHub {
	return &AgentHub{
		Core:             NewCore(),
		start:            make(chan Agent),
		stop:             make(chan Agent),
		registeredAgents: make(map[Agent]bool),
	}
}

func (h *AgentHub) StartAgent(a Agent) {
	h.start <- a
}

func (h *AgentHub) StopAgent(a Agent) {
	h.stop <- a
}

func (h *AgentHub) Run() {
	h.startup()

	stop := *h.Core.StopChannel()

Run:
	for {
		select {
		case a := <-h.start:
			go a.Run()
			h.registeredAgents[a] = true
		case a := <-h.stop:
			go a.Stop()
			delete(h.registeredAgents, a)
		case b := <-stop:
			if b {
				h.shutdown()
				break Run
			}
		}
	}
}

func (h *AgentHub) startup() {
	h.Core.Startup()
}

func (h *AgentHub) shutdown() {
	for a, _ := range h.registeredAgents {
		go a.Stop()
	}

	h.Core.Shutdown()
}
