package main

func SetCube(x, y, z int, instruction bool) bool {
	grid := make(Grid)
	if x < minX || x > maxX || y < minY || y > maxY || z < minZ || z > maxZ {
		// Reject the change
		return false
	}

	// If all coordinates are within range, update the values
	if instruction { // Check if the instruction is to set the cube ON
		if !grid[Coord{x, y, z}] { // Check if the current grid cube is OFF
			cubeCount++ // If the cube is OFF, Turn it ON and raise cubeCount
			grid[Coord{x, y, z}] = true
		}
		return true
	}
	if grid[Coord{x, y, z}] {
		grid[Coord{x, y, z}] = false // If cube is currently ON, turn it OFF and lower cubeCount
		cubeCount--
	}
	return true

}
