package autonomous

import (
	"sync"

	"github.com/elos/data"
)

type Agent interface {
	Run()
	Stop()

	Alive() bool
	Started() *sync.Cond
	Stopped() *sync.Cond

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
