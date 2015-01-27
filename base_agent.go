package autonomous

import (
	"github.com/elos/data"
	"log"
	"sync"
)

func NewBaseAgent() *BaseAgent {
	return &BaseAgent{
		stop: new(chan bool),
	}
}

type BaseAgent struct {
	running bool
	stop    *chan bool

	manager   Manager
	processes int

	m sync.Mutex
}

func (b *BaseAgent) SetManager(m Manager) {
	b.m.Lock()
	defer b.m.Unlock()
	b.manager = m
}

func (b *BaseAgent) Manager() Manager {
	b.m.Lock()
	defer b.m.Unlock()

	return b.manager
}

func (b *BaseAgent) Stop() {
	go func() { *(b.stop) <- true }()
}

func (b *BaseAgent) Kill() {
	// non-blocking
	go func() { *(b.stop) <- true }()
}

func (b *BaseAgent) Alive() bool {
	b.m.Lock()
	defer b.m.Unlock()
	log.Printf("This is the alive function, this agent is: %s", b.running)

	return b.running
}

func (b *BaseAgent) IncrementProcesses() {
	b.m.Lock()
	defer b.m.Unlock()

	b.processes += 1
}

func (b *BaseAgent) DecrementProcesses() {
	b.m.Lock()
	defer b.m.Unlock()

	b.processes -= 1
}

func (b *BaseAgent) Run() {
	b.m.Lock()
	defer b.m.Unlock()

	b.running = true
}

func (b *BaseAgent) StopChannel() *chan bool {
	return b.stop
}

func (b *BaseAgent) Startup() {
	b.m.Lock()
	defer b.m.Unlock()
	b.running = true
}

func (b *BaseAgent) Shutdown() {
	b.m.Lock()
	defer b.m.Unlock()
	b.running = false
}

type BaseDataAgent struct {
	dataAgent data.Identifiable
	m         sync.Mutex
}

func NewBaseDataAgent() *BaseDataAgent {
	return &BaseDataAgent{}
}

func (d *BaseDataAgent) SetDataOwner(a data.Identifiable) {
	d.m.Lock()
	defer d.m.Unlock()
	d.dataAgent = a
}

func (d *BaseDataAgent) DataOwner() data.Identifiable {
	d.m.Lock()
	defer d.m.Unlock()

	return d.dataAgent
}
