package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	ipList := getIpListFromFilename("ip.txt")
	ipListCopy := make([]int, len(ipList))
	copy(ipListCopy, ipList)

	numSteps := getNumJumpsForList(ipList, false)
	fmt.Println("p1 Number of steps - ", numSteps)

	numSteps = getNumJumpsForList(ipListCopy, true)
	fmt.Println("p2 Number of steps - ", numSteps)
}

func getIpListFromFilename(filename string) []int {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")

	ipList := []int{}

	for _, ipByte := range ipStrList {
		ipStr := string(ipByte)
		intVal, strConvErr := strconv.Atoi(ipStr)
		if strConvErr != nil {
			panic(strConvErr)
		}
		ipList = append(ipList, intVal)
	}

	return ipList
}

func getNumJumpsForList(ipList []int, isP2 bool) int {
	numSteps := 0
	currLoc := 0

	for currLoc < len(ipList) && currLoc >= 0 {
		numSteps += 1
		currVal := ipList[currLoc]

		// Solution to p2
		inc := 1
		if isP2 {
			if currVal >= 3 {
				inc = -1
			}
		}

		ipList[currLoc] = currVal + inc
		currLoc = currLoc + currVal
	}
	return numSteps
}
