package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MrG00gle/Simple-Sheduler/sheduler"
)

func generateTimestamps(numTimestamps int, minDuration uint32, maxSpread uint32) (timestamps []uint32) {
	timestamps = make([]uint32, numTimestamps)
	timestamps[0] = 0

	for i := 1; i < numTimestamps; i++ {
		duration := uint32(rand.Intn(int(maxSpread))) + minDuration
		newTimestamp := timestamps[i-1] + duration
		if newTimestamp < timestamps[i-1] { // Overflow check
			fmt.Printf("Warning: Timestamp overflow at index %d, capping at max uint32\n", i)
			newTimestamp = 4294967295 // Max uint32
		}
		timestamps[i] = newTimestamp
	}
	return timestamps
}

func generateTasks(taskNumber uint16, timestamps []uint32, operation sheduler.Operation) []*sheduler.Task {
	tasks := make([]*sheduler.Task, taskNumber)
	for i := 0; i < int(taskNumber); i++ {
		tasks[i] = &sheduler.Task{ID: uint16(i), Stamps: timestamps, Operation: operation}
	}
	return tasks
}

func oper(index int, id uint16, status sheduler.TaskStatus, stamp uint32) {
	fmt.Printf("Index: %d, ID: %d, Status: %v, Stamp: %d\n", index, id, status, stamp)
}

func main() {
	// Seed random number generator for reproducibility
	rand.Seed(time.Now().UnixNano())

	shed := sheduler.Sheduler{}
	timestamps := generateTimestamps(100, 100, 2000)
	operation := sheduler.Operation(oper)
	tasks := generateTasks(10, timestamps, operation)

	for _, task := range tasks {
		shed.AddThread(task)
	}

	shed.Start()

	//go shed.Start()

	//Example: Pause and resume the first task
	//if len(tasks) > 0 {
	//	time.Sleep(500 * time.Millisecond) // Wait for some tasks to start
	//	shed.PauseThread(tasks[0])
	//	fmt.Println("Paused task 0")
	//	time.Sleep(500 * time.Millisecond)
	//	shed.ResumeThread(tasks[0])
	//	fmt.Println("Resumed task 0")
	//}

	// Wait for all tasks to complete (optional, for demonstration)
	//time.Sleep(10 * time.Second)
}
