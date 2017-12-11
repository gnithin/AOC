package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func findNumSteps(positions []string) (int, int) {
	maxDistance := -1
	finalPos := []int{0, 0, 0}
	for _, pos := range positions {
		x, y, z := posToCoordMap(pos)
		finalPos[0] += x
		finalPos[1] += y
		finalPos[2] += z

		currDistance := calcDistanceFromOrigin(finalPos)
		if currDistance > maxDistance {
			maxDistance = currDistance
		}
	}

	// Find the max
	return calcDistanceFromOrigin(finalPos), maxDistance
}

// This calculation is from this awesome page -
// https://www.redblobgames.com/grids/hexagons/
// I went to the point of figuring out that there would be 3 axes
// (4 directions, hence 2 axes, 6 directions so 3 axes :p), but could
// not go further. This page beautifully explains that.
func calcDistanceFromOrigin(pos []int) int {
	maxVal := 0
	for _, val := range pos {
		maxVal += int(math.Abs(float64(val)))
	}
	return maxVal / 2
}

func posToCoordMap(position string) (int, int, int) {
	switch position {
	case "n":
		return 0, 1, -1
	case "s":
		return 0, -1, 1
	case "ne":
		return 1, 0, -1
	case "nw":
		return -1, 1, 0
	case "se":
		return 1, -1, 0
	case "sw":
		return -1, 0, 1
	}
	return 0, 0, 0
}

func main() {
	filename := "ip.txt"
	ipList := getIpListFromFilename(filename)
	for _, ip := range ipList {
		shortestNumSteps, maxDistance := findNumSteps(ip)
		fmt.Println("Shortest number of steps - ", shortestNumSteps)
		fmt.Println("Max distance - ", maxDistance)
	}
}

func getIpListFromFilename(filename string) [][]string {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	var ipCompList [][]string
	for _, ipStr := range ipStrList {
		ipComp := strings.Split(ipStr, ",")
		ipCompList = append(ipCompList, ipComp)
	}
	return ipCompList
}
