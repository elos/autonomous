package autonomous

type AgentHub struct {
	*BaseAgent

	start            chan Agent
	stop             chan Agent
	registeredAgents map[Agent]bool
}

func NewAgentHub() *AgentHub {
	return &AgentHub{
		BaseAgent:        NewBaseAgent(),
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
	stop := h.BaseAgent.StopChannel()

	for {
		select {
		case a := <-h.start:
			go a.Run()
			h.registeredAgents[a] = true
		case a := <-h.stop:
			go a.Run()
			delete(h.registeredAgents, a)
		case _ = <-*stop:
			h.shutdown()
			break
		}
	}
}

func (h *AgentHub) startup() {
	h.BaseAgent.Startup()
}

func (h *AgentHub) shutdown() {
	for a, _ := range h.registeredAgents {
		go a.Stop()
	}

	h.BaseAgent.Shutdown()
}
