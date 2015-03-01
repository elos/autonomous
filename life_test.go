package autonomous

import (
	"testing"
	"time"
)

func TestLife(t *testing.T) {

	var testLife = NewLife()

	if testLife == nil {
		t.Errorf("test life object is nil")
	} else {
		t.Log("Recieved a non-nil Life object")
	}

	if testLife.Alive() != false {
		t.Errorf("a life should start as not alive")
	} else {
		t.Log("Test object starts with alive = false")
	}

	c := make(chan bool)

	go func() {
		t.Log("Waiting for start")
		testLife.WaitStart()
		c <- true
		t.Log("Exited from WaitStart")
	}()

	go func() {
		t.Log("Beginning testLife - should be after 'Waiting for Start'")
		testLife.Begin()
	}()

	select {
	case <-time.After(5 * time.Second):
		t.Errorf("Waiting for testlife.WaitStart() timed out")
	case <-c:
		t.Log("Successfully WaitStarted")
	}

	if testLife.Alive() != true {
		t.Errorf("Begin() should start the life, and the life should be alive")
	} else {
		t.Log("Begin succesfully changed alive to true")
	}

	go func() {
		t.Log("Waiting for Stop")
		testLife.WaitStop()
		c <- true
		t.Log("Exited from WaitStop")
	}()

	go func() {
		testLife.WaitStop()
	}()

	go func() {
		t.Log("Ending testLife - should be after 'Waiting for Stop'")
		testLife.End()
	}()

	select {
	case <-time.After(5 * time.Second):
		t.Errorf("Waiting for testLife.WaitStop() timed out")
	case <-c:
		t.Log("Succesfully WaitStopped")
	}

	if testLife.Alive() != false {
		t.Errorf("End() should end the life, and alive should be false")
	} else {
		t.Log("End changed alive to false")
	}

	// Mock concurrency?

	i := 0

	for i > 5 {
		go func() {
			testLife.Begin()
		}()

		i++
	}

	for i > 0 {
		go func() {
			testLife.End()
		}()

		i--
	}

	testLife.End()

	if testLife.Alive() != false {
		t.Errorf("concurrent access somehow messed with life")
	}

	multi := make(chan int)
	i = 0
	for i < 3 {
		go func() {
			testLife.WaitStart()
			multi <- 1
		}()

		i++
	}

	go func() {
		testLife.Begin()
	}()

	i = 3
	for i > 0 {
		select {
		case <-time.After(1 * time.Second):
			t.Errorf("Timeout on multi WaitStarts")
		case <-multi:
			t.Log("WaitStart Received")
			i--
		}
	}

	t.Log("Successfully handled multi WaitStarts")
}

func TestMultipleLives(t *testing.T) {
	testLife := NewLife()
	testLife.Begin()

	c := make(chan bool)
	go func() {
		testLife.WaitStop()
		c <- true
	}()

	go testLife.End()

	select {
	case <-time.After(1 * time.Second):
		t.Errorf("Timeout on first life wait start")
	case <-c:
	}

	testLife.Begin()

	if testLife.Alive() != true {
		t.Errorf("Test life should alive in its second life")
	}

	c = make(chan bool)
	go func() {
		testLife.WaitStop()
		c <- true
	}()

	go testLife.End()
	testLife.WaitStop()

	select {
	case <-time.After(1 * time.Second):
		t.Errorf("Timeout on second life wait start")
	case <-c:
	}

	testLife.Begin()
	go testLife.End()
	testLife.WaitStop()
	if testLife.Alive() != false {
		t.Errorf("testLife should be dead")
	}

}

func TestMultipleWaiters(t *testing.T) {
	testLife := NewLife()

	go testLife.Begin()
	testLife.WaitStart()
	if testLife.Alive() != true {
		t.Errorf("testlife should be alive")
	}

	one := make(chan bool)
	two := make(chan bool)
	three := make(chan bool)

	go func() {
		testLife.WaitStop()
		one <- true
	}()

	go func() {
		testLife.WaitStop()
		two <- true
	}()

	go func() {
		testLife.WaitStop()
		three <- true
	}()

	go testLife.End()

	select {
	case <-time.After(1 * time.Second):
		t.Errorf("Timeout on first WaitStop")
	case <-one:
	}

	select {
	case <-time.After(1 * time.Second):
		t.Errorf("Timeout on second second WaitStop")
	case <-two:
	}

	select {
	case <-time.After(1 * time.Second):
		t.Errorf("Timeout on third WaitStop")
	case <-three:
	}

	final := make(chan bool)
	go func() {
		testLife.WaitStop()
		final <- true
	}()

	select {
	case <-time.After(20 * time.Millisecond):
	case <-final:
		t.Errorf("Should not have exited WaitStop")
	}

	testLife.End()
	<-final
}

func TestWaitStartLocking(t *testing.T) {
}
