package main

import "fmt"

const filename string = "instructions.txt"

func main() {
	// var ranges coordRange
	// r := make(map[int]coordRange)

	lines := readFile() // read reboot instructions from file

	for i := 0; i < len(lines); i++ {
		// r[i] = getCoordRanges(lines[i]) // retrieve usable ranges of x,y,z coords
		coords := getCoordRanges(lines[i])
		fmt.Printf("Working on line: %v\n", lines[i])
		for xMin := coords.x[0]; xMin < coords.x[1]; xMin++ {
			for yMin := coords.y[0]; yMin <= coords.y[1]; yMin++ {
				for zMin := coords.z[0]; zMin <= coords.z[1]; zMin++ {
					if !SetCube(xMin, yMin, zMin, coords.on) {
						// fmt.Printf("Unable to set cube:\t%d.%d.%d-%v", xMin, yMin, zMin, coords.on)
						continue
					}
				}
			}
		}
	}
	fmt.Printf("# of cubes turned on: %d", cubeCount)
}
