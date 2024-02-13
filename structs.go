package main

type coordRange struct {
	on bool
	// [0] is start of range, [1] the end of range
	x [2]int
	y [2]int
	z [2]int
}

const (
	minX = -50
	maxX = 50
	minY = -50
	maxY = 50
	minZ = -50
	maxZ = 50
)

var cubeCount int

type Coord struct {
	x, y, z int
}

type Grid map[Coord]bool
