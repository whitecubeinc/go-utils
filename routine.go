package utils

import (
	"runtime/debug"
	"sync"
	"time"
)

func GoroutineWithTicker(routine func(), duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		defer func() {
			ticker.Stop()
			if p := recover(); p != nil {
				debugStack := debug.Stack()
				Error(p)
				Error(string(debugStack))
			}
		}()
		for {
			<-ticker.C
			routine()
		}
	}()
}

func GoroutineWithTimer(wg *sync.WaitGroup, routine func(), duration time.Duration) {
	if wg != nil {
		wg.Add(1)
	}

	timer := time.NewTimer(duration)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				debugStack := debug.Stack()
				Error(p)
				Error(string(debugStack))
			}
			if wg != nil {
				wg.Done()
			}
		}()
		<-timer.C
		routine()
	}()
}

func GoroutineLimitedNumber(wg *sync.WaitGroup, routineQ *chan struct{}, routine func()) {
	wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				debugStack := debug.Stack()
				Error(p)
				Error(string(debugStack))
			}

			wg.Done()
		}()
		*routineQ <- struct{}{}
		routine()
		<-*routineQ
	}()
}

func GoroutineWithWaitGroup(wg *sync.WaitGroup, routine func()) {
	wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				debugStack := debug.Stack()
				Error(p)
				Error(string(debugStack))
			}

			wg.Done()
		}()
		routine()
	}()
}
