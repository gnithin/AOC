package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	ipList := getIpListFromFilename("ip.txt")
	numSteps := getNumJumpsForList(ipList)
	fmt.Println("Number of steps - ", numSteps)
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

func getNumJumpsForList(ipList []int) int {
	numSteps := 0
	currLoc := 0

	for currLoc < len(ipList) && currLoc >= 0 {
		numSteps += 1
		currVal := ipList[currLoc]
		ipList[currLoc] = currVal + 1
		currLoc = currLoc + currVal
	}
	return numSteps
}
