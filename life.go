package autonomous

import "sync"

type Life interface {
	Alive() bool
	Begin()
	End()
	Started() *sync.Cond
	Stopped() *sync.Cond
}

type life struct {
	alive   bool
	m       *sync.Mutex
	started *sync.Cond
	stopped *sync.Cond
	stop    chan bool
}

func NewLife() Life {
	l := &life{
		alive: false,
		m:     new(sync.Mutex),
		stop:  make(chan bool),
	}

	l.started = sync.NewCond(l.m)
	l.stopped = sync.NewCond(l.m)

	return l
}

func (l *life) Begin() {
	l.m.Lock()
	defer l.m.Unlock()

	l.alive = true
	l.started.Broadcast()
}

func (l *life) End() {
	l.m.Lock()
	defer l.m.Unlock()

	l.alive = false
	l.stopped.Broadcast()
}

func (l *life) Started() *sync.Cond {
	return l.started
}

func (l *life) Stopped() *sync.Cond {
	return l.stopped
}

func (l *life) Alive() bool {
	l.m.Lock()
	defer l.m.Unlock()

	return l.alive
}
