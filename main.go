package main

import "fmt"

const filename string = "initSteps.txt"

func main() {
	var ranges coordRange
	lines := readFile() // read reboot instructions from file

	for i := 0; i < len(lines); i++ {
		ranges = getCoordRanges(lines[i])
	}
	fmt.Printf("X ranges:\t%d - %d\n", ranges.x[0], ranges.x[1])
}
