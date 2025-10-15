package sheduler

type Task struct {
	ID     uint16
	Status TaskStatus
	Stamps []uint32
}

type TaskStatus uint8

const (
	StatusPending TaskStatus = iota
	StatusRunning
	StatusPaused
	StatusDone
	StatusAborted
)
