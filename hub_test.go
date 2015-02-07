package autonomous_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	. "github.com/elos/autonomous"
)

type TestAgent struct {
	Life
	Managed
}

func NewTestAgent() Agent {
	return &TestAgent{
		Life: NewLife(),
	}
}

func (t *TestAgent) Start() {
	t.Life.Begin()
}

func (t *TestAgent) Stop() {
	t.Life.End()
}

func TestNewHub(t *testing.T) {
	h := NewHub()

	if h == nil {
		t.Errorf("NewHub returned nil")
	}
}

func TestHub(t *testing.T) {
	h := NewHub()

	go h.Start()

	ta := NewTestAgent()

	c := make(chan bool)
	go func() {
		ta.WaitStart()
		c <- true
	}()

	h.StartAgent(ta)

	select {
	case <-time.After(50 * time.Millisecond):
		t.Errorf("Timed out waiting for agent to start")
	case <-c:
		// ready to check alive
	}

	if ta.Alive() != true {
		t.Errorf("Hub didn't start test agent")
	}

	_, ok := h.Agents()[ta]

	if !ok {
		t.Errorf("Hub doesn't know about the test agent")
	}
	go func() {
		ta.WaitStop()
		c <- true
	}()
	h.StopAgent(ta)

	select {
	case <-time.After(50 * time.Millisecond):
		t.Errorf("Timed out waiting for agent to stop")
	case <-c:
		// ready to check alive
	}

	if ta.Alive() != false {
		t.Errorf("Hub didn't stop test agent")
	}

	_, ok = h.Agents()[ta]

	if ok {
		t.Errorf("Hub still knows about test agent")
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			ta := NewTestAgent()

			h.StartAgent(ta)
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	if len(h.Agents()) != 100 {
		t.Errorf("Not all agents made it to hub")
	}

	for agent, _ := range h.Agents() {
		if agent.Alive() != true {
			t.Errorf("One agent wasn't correctly started")
		}

		h.StopAgent(agent)
	}

	if len(h.Agents()) != 0 {
		t.Errorf("Hub failed to stop all agents")
	}

	if h.Alive() != true {
		t.Errorf("Hub should still be alive")
	}

	// now we can assume StartAgent works
	h.StartAgent(ta)

	go func() {
		h.WaitStop()
		c <- true
	}()

	h.Stop()

	select {
	case <-time.After(50 * time.Millisecond):
		t.Errorf("Timed out waiting for hub to stop")
	case <-c:
		// good to go
	}

	if ta.Alive() == true {
		t.Errorf("Hub Stop should kill all agents (and wait for their deaths)")
	}

	if h.Alive() == true {
		t.Errorf("Hub should be dead")
	}
}
