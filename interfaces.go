package autonomous

type (
	Agent interface {
		Start()
		Stop()

		Alive() bool
		WaitStart()
		WaitStop()

		SetManager(Manager)
		Manager() Manager
	}

	Manager interface {
		Agent

		StartAgent(Agent)
		StopAgent(Agent)
	}
)
