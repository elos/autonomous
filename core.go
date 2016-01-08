package autonomous

import (
	"sync"
)

// --- Managed {{{

type Managed struct {
	manager Manager
	sync.Mutex
}

func (m *Managed) Manager() Manager {
	m.Lock()
	defer m.Unlock()

	return m.manager
}

func (m *Managed) SetManager(newM Manager) {
	m.Lock()
	defer m.Unlock()

	m.manager = newM
}

// --- }}

// --- Tallied {{{

type Tallied struct {
	tally int
	sync.Mutex
}

func (t *Tallied) Tally() int {
	t.Lock()
	defer t.Unlock()

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

// --- }}}
