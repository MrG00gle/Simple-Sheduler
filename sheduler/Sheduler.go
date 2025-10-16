package sheduler

import (
	"sync"
	"time"
)

type Sheduler struct {
	tasks []Task
}

func (sheduler *Sheduler) AddThread(task Task) {
	task.Status = StatusPending
	sheduler.tasks = append(sheduler.tasks, task)
}

func (sheduler *Sheduler) PauseThread(task Task) {
	task.Status = StatusPaused
}

func (sheduler *Sheduler) ResumeThread(task Task) {
	task.Status = StatusRunning
}

func (sheduler *Sheduler) AbortThread(task Task) {
	task.Status = StatusAborted
}

func (sheduler *Sheduler) Thread(task Task) {
	var paststamp uint32 = 0
	task.Status = StatusRunning
	for index, stamp := range task.Stamps {
		time.Sleep(time.Duration(stamp-paststamp) * time.Millisecond)
		paststamp = stamp

		switch task.Status {
		case StatusRunning:
			task.Operation(index, task.ID, task.Status, stamp)
		case StatusPaused:
			continue
		case StatusAborted:
			break
		case StatusDone:
			break
		default:
			panic("unhandled default case")
		}

		if index == len(sheduler.tasks) {
			task.Status = StatusDone
		}

		//if task.Status == StatusPaused {
		//	continue
		//} else if task.Status == StatusRunning {
		//	task.Operation(index, task.ID, task.Status, stamp)
		//}
	}
}

func (sheduler *Sheduler) Start() {
	var wg sync.WaitGroup
	wg.Add(len(sheduler.tasks))
	for _, task := range sheduler.tasks {
		go sheduler.Thread(task)
	}
	wg.Wait()
}
