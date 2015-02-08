package autonomous

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestManaged(t *testing.T) {
	m := new(Managed)

	if m == nil {
		t.Errorf("new(Managed) should not return nil")
	}

	var h = NewHub()

	m.SetManager(h)

	if m.Manager() != h {
		t.Errorf("SetManager failed to set the manager expected %+v, got: %+v", h, m.Manager())
	}

	i := 0
	for i < 10 {
		one := make(chan bool)
		two := make(chan Manager)
		hPrime := NewHub()

		go func() {
			m.SetManager(hPrime)
			one <- true
		}()

		go func() {
			two <- m.Manager()
		}()

		select {
		case <-one:
			if m.Manager() != hPrime {
				t.Errorf("m.Manager() should be hPrime")
			}

			<-two
		case fo := <-two:
			if fo != h {
				t.Error("m.Manager() should be h")
			}

			<-one
		}

		i++
	}
}

func TestTallied(t *testing.T) {
	tally := new(Tallied)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

			if i%2 == 0 {
				if i%5 == 0 {
					tally.Drop(5)
				} else {
					tally.Decr()
				}
			} else {
				if i%5 == 0 {
					tally.Add(5)
				} else {
					tally.Incr()
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	if tally.Tally() != 0 {
		t.Errorf("Tally is not thread-safe expected 0, got %d", tally.Tally())
	}

}

func TestStopper(t *testing.T) {
	s := make(Stopper)

	go s.Stop()

	select {
	case <-s:
		t.Log("Stopper Stop() passed bool")
	case <-time.After(50 * time.Millisecond):
		t.Errorf("Timed out wating for stopper")
	}
}
