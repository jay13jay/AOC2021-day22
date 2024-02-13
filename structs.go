package main

import (
	"sync"
)

type coordRange struct {
	on bool
	// [0] is start of range, [1] the end of range
	x [2]int
	y [2]int
	z [2]int
}

const (
	minX              = -50
	maxX              = 50
	minY              = -50
	maxY              = 50
	minZ              = -50
	maxZ              = 50
	filename   string = "instructions.txt"
	numWorkers        = 100
)

var cubeCount int
var mutex sync.Mutex          // Mutex to protect access to the cubeCount var
var taskQueueMutex sync.Mutex // Mutex to protect access to taskQueues slice

type Coord struct {
	x, y, z int
}

type Grid map[Coord]bool
