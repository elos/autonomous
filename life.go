package autonomous

import "sync"

type Life interface {
	Alive() bool
	Begin()
	End()
	WaitStart()
	WaitStop()
}

type life struct {
	alive bool
	m     *sync.Mutex
	bm    *sync.Mutex
	em    *sync.Mutex
	stop  chan bool

	started *sync.Cond
	stopped *sync.Cond
}

func NewLife() Life {
	l := &life{
		alive: false,
		m:     new(sync.Mutex),
		em:    new(sync.Mutex),
		bm:    new(sync.Mutex),
		stop:  make(chan bool),
	}

	l.started = sync.NewCond(l.em)
	l.stopped = sync.NewCond(l.bm)

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

func (l *life) Alive() bool {
	l.m.Lock()
	defer l.m.Unlock()

	return l.alive
}

func (l *life) WaitStart() {
	l.started.L.Lock()
	l.started.Wait()
	l.started.L.Unlock()
}

func (l *life) WaitStop() {
	l.stopped.L.Lock()
	l.stopped.Wait()
	l.stopped.L.Unlock()
}
