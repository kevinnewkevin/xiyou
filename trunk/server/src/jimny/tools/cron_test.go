package tools

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tk := NewCron("taska", "0/30 * * * * *", func() error { fmt.Println("hello world"); return nil })
	err := tk.Run()
	if err != nil {
		t.Fatal(err)
	}
	AddCron("taska", tk)
	StartCron()
	time.Sleep(6 * time.Second)
	StopCron()
}

func TestSpec(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	tk1 := NewCron("tk1", "0 12 * * * *", func() error { fmt.Println("tk1"); return nil })
	tk2 := NewCron("tk2", "0,10,20 * * * * *", func() error { fmt.Println("tk2"); wg.Done(); return nil })
	tk3 := NewCron("tk3", "0 10 * * * *", func() error { fmt.Println("tk3"); wg.Done(); return nil })

	AddCron("tk1", tk1)
	AddCron("tk2", tk2)
	AddCron("tk3", tk3)

	StartCron()

	defer StopCron()

	select {
	case <-time.After(200 * time.Second):
		t.FailNow()
	case <-wait(wg):
	}
}

func wait(wg *sync.WaitGroup) chan bool {
	ch := make(chan bool)
	go func() {
		wg.Wait()
		ch <- true
	}()
	return ch
}
