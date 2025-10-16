package sheduler

import (
	"sync"
	"time"
)

type Sheduler struct {
	tasks []*Task // Store pointers to tasks
}

func (sheduler *Sheduler) AddThread(task *Task) {
	task.mu.Lock()
	task.Status = StatusPending
	task.mu.Unlock()
	sheduler.tasks = append(sheduler.tasks, task)
}

func (sheduler *Sheduler) PauseThread(task *Task) {
	task.mu.Lock()
	task.Status = StatusPaused
	task.mu.Unlock()
}

func (sheduler *Sheduler) ResumeThread(task *Task) {
	task.mu.Lock()
	task.Status = StatusRunning
	task.mu.Unlock()
}

func (sheduler *Sheduler) AbortThread(task *Task) {
	task.mu.Lock()
	task.Status = StatusAborted
	task.mu.Unlock()
}

func (sheduler *Sheduler) Thread(task *Task, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure wg.Done() is called when goroutine exits

	var paststamp uint32 = 0
	task.mu.Lock()
	task.Status = StatusRunning
	task.mu.Unlock()

	for index, stamp := range task.Stamps {
		time.Sleep(time.Duration(stamp-paststamp) * time.Millisecond)
		paststamp = stamp

		task.mu.RLock()
		currentStatus := task.Status
		task.mu.RUnlock()

		switch currentStatus {
		case StatusRunning:
			task.Operation(index, task.ID, currentStatus, stamp)
		case StatusPaused:
			continue
		case StatusAborted, StatusDone:
			return
		default:
			panic("unhandled default case")
		}

		// Mark task as done after the last stamp
		if index == len(task.Stamps)-1 {
			task.mu.Lock()
			task.Status = StatusDone
			task.mu.Unlock()
			return
		}
	}
}

func (sheduler *Sheduler) Start() {
	var wg sync.WaitGroup
	wg.Add(len(sheduler.tasks))
	for _, task := range sheduler.tasks {
		go sheduler.Thread(task, &wg) // Pass wg to Thread
	}
	wg.Wait()
}
