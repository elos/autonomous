package autonomous

type Agent interface {
	Start()
	Stop()

	Alive() bool
	WaitStart()
	WaitStop()

	SetManager(Manager)
	Manager() Manager
}

type Manager interface {
	Agent

	StartAgent(Agent)
	StopAgent(Agent)
}
