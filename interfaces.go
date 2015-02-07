package autonomous

import (
	"github.com/elos/data"
)

type Agent interface {
	Start()
	Stop()

	Alive() bool
	WaitStart()
	WaitStop()

	SetManager(Manager)
	Manager() Manager
}

type DataAgent interface {
	SetDataOwner(data.Identifiable)
	DataOwner() data.Identifiable
}

type Manager interface {
	Agent

	StartAgent(Agent)
	StopAgent(Agent)
}
