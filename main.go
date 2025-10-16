package main

import (
	"fmt"
	"math/rand"

	"github.com/MrG00gle/Simple-Sheduler/sheduler"
)

func generateTimestamps(numTimestamps int, minDuration uint32, maxSpread uint32) (timestamps []uint32) {
	// Number of timestamps to generate
	timestamps = make([]uint32, numTimestamps)

	// Set starting timestamp to 0
	timestamps[0] = 0

	// Generate timestamps with random durations
	// Adjust range to fit within uint32 (0 to 4,294,967,295)
	//minDuration := uint32(100)    // Minimum duration in milliseconds
	//maxSpread := uint32(1000)     // Maximum additional spread (e.g., 100 to 1100 ms
	for i := 1; i < numTimestamps; i++ {
		// Generate random duration
		duration := uint32(rand.Intn(int(maxSpread))) + minDuration
		newTimestamp := timestamps[i-1] + duration
		if newTimestamp < timestamps[i-1] { // Overflow chek
			newTimestamp = 4294967295 // Set Cap of uint32
		}
		timestamps[i] = newTimestamp
	}

	return timestamps
}

func generateTasks(tasknumber uint16, timestamps []uint32, operation sheduler.Operation) []sheduler.Task {
	tasks := make([]sheduler.Task, tasknumber)
	for i := 0; i < int(tasknumber); i++ {
		tasks[i] = sheduler.Task{ID: uint16(i), Stamps: timestamps, Operation: operation}
	}
	return tasks
}

func oper(index int, id uint16, status sheduler.TaskStatus, stamp uint32) {
	fmt.Println("Index:", index, "ID", id, "Status", status, "Stamp", stamp)
}

func main() {
	shed := sheduler.Sheduler{}
	timestamps := generateTimestamps(50, 100, 1000)
	operation := sheduler.Operation(oper)
	tasks := generateTasks(10, timestamps, operation)

	for _, task := range tasks {
		shed.AddThread(task)
	}

	shed.Start()
}
