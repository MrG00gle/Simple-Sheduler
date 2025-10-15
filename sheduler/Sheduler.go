package sheduler

import "time"

type Sheduler struct {
	tasks []Task
}

func (sheduler *Sheduler) AddTask(task Task) {
	sheduler.tasks = append(sheduler.tasks, task)
}

func (sheduler *Sheduler) Thread(task Task) {
	var paststamp uint32 = 0
	for id, stamp := range task.Stamps {
		// do something here than wait
		time.Sleep(time.Duration(stamp-paststamp) * time.Millisecond)
		if task.Status == StatusPaused {
			continue
		}
	}

}
