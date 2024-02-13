package main

import (
	"log"
	"strconv"
	"strings"
)

func getCoordRanges(line string) coordRange {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '=' // replace all the symbols we don't need with spaces
	})
	// fmt.Printf("Whole Line: %v\n", fields) // Debug line

	var coords coordRange

	instruction := fields[0]

	if instruction == "on" {
		coords.on = true
	} else if instruction == "off" {
		coords.on = false
	} else {
		log.Fatal("Unable to process instruction")
	}

	/*
		After split, a line will look like the following:
		FIELD |0 |1| 2| 3| 4| 5 |6 |7| 8| 9 |
		Line: [on x -20 26 y -36 17 z -47 7]
	*/

	// Get x ranges
	x1, _ := strconv.Atoi(fields[2])
	x2, _ := strconv.Atoi(fields[3])
	coords.x = [2]int{
		x1,
		x2,
	}

	y1, _ := strconv.Atoi(fields[5])
	y2, _ := strconv.Atoi(fields[6])
	coords.y = [2]int{
		y1,
		y2,
	}

	z1, _ := strconv.Atoi(fields[8])
	z2, _ := strconv.Atoi(fields[9])
	coords.z = [2]int{
		z1,
		z2,
	}
	// fmt.Printf("Coords var:\t %v\n", coords)
	return coords

}
