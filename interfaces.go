package autonomous

import (
	"time"

	"sync"

	"github.com/elos/data"
)

type Agent interface {
	Run()
	Stop()
	Kill()
	Alive() *sync.Cond

	SetManager(Manager)
	Manager() Manager
}

type DataAgent interface {
	SetDataOwner(data.Identifiable)
	DataOwner() data.Identifiable
}

type NewAgent func(db data.DB, a data.Identifiable, d time.Duration) Agent

type Manager interface {
	Agent

	StartAgent(Agent)
	StopAgent(Agent)
}
