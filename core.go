package autonomous

import (
	"github.com/elos/data"
	"sync"
)

type Managed struct {
	manager Manager
	sync.RWMutex
}

func (m *Managed) Manager() Manager {
	m.RLock()
	defer m.RUnlock()

	return m.manager
}

func (m *Managed) SetManager(man Manager) {
	m.Lock()
	defer m.Unlock()
	m.manager = man
}

type Tallied struct {
	tally int
	sync.RWMutex
}

func (t *Tallied) Tally() int {
	t.RLock()
	defer t.RUnlock()

	return t.tally
}

func (t *Tallied) Incr() {
	t.Add(1)
}

func (t *Tallied) Decr() {
	t.Drop(1)
}

func (t *Tallied) Add(delta int) {
	t.Lock()
	defer t.Unlock()

	t.tally += delta
}

func (t *Tallied) Drop(delta int) {
	t.Lock()
	defer t.Unlock()

	t.tally -= delta
}

type Stopper chan bool

func (s Stopper) Stop() {
	s <- true
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
