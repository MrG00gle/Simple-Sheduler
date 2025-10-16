package sheduler

import "sync"

type TaskStatus uint8

const (
	StatusPending TaskStatus = iota
	StatusRunning
	StatusPaused
	StatusDone
	StatusAborted
)

type Operation func(index int, id uint16, status TaskStatus, stamp uint32)

type Task struct {
	ID        uint16
	Status    TaskStatus
	Stamps    []uint32
	Operation Operation
	mu        sync.RWMutex // Mutex for thread-safe status access
}
