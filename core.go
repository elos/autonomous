package autonomous

import (
	"github.com/elos/data"
	"log"
	"sync"
)

func NewCore() *Core {
	s := make(chan bool)
	c := &Core{
		stop: &s,
		m:    &sync.Mutex{},
	}

	c.alive = sync.NewCond(c.m)

	return c
}

type Core struct {
	running bool
	alive   *sync.Cond
	stop    *chan bool

	manager   Manager
	processes int

	m *sync.Mutex
}

func (b *Core) SetManager(m Manager) {
	b.m.Lock()
	defer b.m.Unlock()

	b.manager = m
}

func (b *Core) Manager() Manager {
	b.m.Lock()
	defer b.m.Unlock()

	return b.manager
}

func (b *Core) Stop() {
	*(b.stop) <- true
}

func (b *Core) Kill() {
	*(b.stop) <- true
}

func (b *Core) Alive() *sync.Cond {
	b.m.Lock()
	defer b.m.Unlock()
	log.Printf("This is the alive function, this agent is: %s", b.running)

	return b.alive
}

func (b *Core) IncrementProcesses() {
	b.m.Lock()
	defer b.m.Unlock()

	b.processes += 1
}

func (b *Core) DecrementProcesses() {
	b.m.Lock()
	defer b.m.Unlock()

	b.processes -= 1
}

func (b *Core) StopChannel() *chan bool {
	return b.stop
}

func (b *Core) Run() {
	b.Startup()
}

func (b *Core) Startup() {
	b.m.Lock()
	defer b.m.Unlock()
	b.running = true
	b.alive.Broadcast()
}

func (b *Core) Shutdown() {
	b.m.Lock()
	defer b.m.Unlock()
	b.running = false
	b.alive.Broadcast()
}

type Identified struct {
	dataAgent data.Identifiable
	m         sync.Mutex
}

func NewIdentified() *Identified {
	return &Identified{}
}

func (d *Identified) SetDataOwner(a data.Identifiable) {
	d.m.Lock()
	defer d.m.Unlock()
	d.dataAgent = a
}

func (d *Identified) DataOwner() data.Identifiable {
	d.m.Lock()
	defer d.m.Unlock()

	return d.dataAgent
}
