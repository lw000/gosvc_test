package service2

import (
	"log"
	"sync"
	"time"
)

type TimerService struct {
	exit chan struct{}
	wg   sync.WaitGroup
}

func (ts *TimerService) Start() {
	ts.exit = make(chan struct{}, 1)
	ts.wg.Add(1)
	go func() {
		defer func() {
			if x := recover(); x != nil {
			}
			ts.wg.Done()
		}()

		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case d := <-t.C:
				log.Println(1, d.Format("2006-01-02 15:04:05"))
			case <-ts.exit:
				log.Println("EXIT")
				return
			}
		}
	}()
}

func (ts *TimerService) Stop() error {
	close(ts.exit)
	ts.wg.Wait()
	return nil
}
