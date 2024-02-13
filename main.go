package main

import (
	"fmt"
	"sync"
)

func main() {
	lines := readFile() // read reboot instructions from file
	grid := make(Grid)

	var wg sync.WaitGroup

	// Start worker goroutines
	taskQueues := make([]chan Coord, len(lines)) // Slice to hold task queues for each line
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&grid, &wg, taskQueues)
	}

	for i, line := range lines {
		taskQueue := make(chan Coord, 100) // Buffered channel for task coordination
		taskQueueMutex.Lock()
		taskQueues[i] = taskQueue // Assign the task queue for this line to the slice
		taskQueueMutex.Unlock()
		coords := getCoordRanges(line)
		fmt.Printf("Working on line: %v\n", line)
		for xMin := coords.x[0]; xMin < coords.x[1]; xMin++ {
			for yMin := coords.y[0]; yMin <= coords.y[1]; yMin++ {
				for zMin := coords.z[0]; zMin <= coords.z[1]; zMin++ {
					if taskQueue != nil { // Ensure taskQueue is not nil before sending data into it
						taskQueue <- Coord{xMin, yMin, zMin}
					}
				}
			}
		}
		close(taskQueue) // Close the task queue for this line after all tasks are queued
	}

	wg.Wait() // Wait for all workers to finish processing tasks for all lines

	fmt.Printf("# of cubes turned on: %d", cubeCount)
}

func worker(grid *Grid, wg *sync.WaitGroup, taskQueues []chan Coord) {
	defer wg.Done()
	for {
		taskQueueMutex.Lock()
		for _, taskQueue := range taskQueues {
			if taskQueue == nil { // Skip nil channels
				continue
			}
			for coord := range taskQueue {
				if !SetCube(grid, coord.x, coord.y, coord.z, true) {
					// fmt.Printf("Unable to set cube:\t%d.%d.%d-%v", coord.x, coord.y, coord.z, coords.on)
					continue
				}
			}
		}
		taskQueueMutex.Unlock()
	}
}

func SetCube(grid *Grid, x, y, z int, on bool) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if x < minX || x > maxX || y < minY || y > maxY || z < minZ || z > maxZ {
		// Reject the change
		return false
	}

	// If all coordinates are within range, update the values
	if on { // Check if the instruction is to set the cube ON
		if !(*grid)[Coord{x, y, z}] { // Check if the current grid cube is OFF
			cubeCount++ // If the cube is OFF, Turn it ON and raise cubeCount
			(*grid)[Coord{x, y, z}] = true
		}
		return true
	}

	if (*grid)[Coord{x, y, z}] {
		(*grid)[Coord{x, y, z}] = false // If cube is currently ON, turn it OFF and lower cubeCount
		cubeCount--
	}
	return true
}
