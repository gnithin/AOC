package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	gridMap := getIpListFromFilename(filename)
	grid := createGridFromInitialMap(gridMap)

	// P2 totally overrides P1. Please look at commit - 289c5e for P1
	burstSize := 10000000
	virus := createVirusWithGrid(&grid)
	virus.infectWithBurstSize(burstSize)
	fmt.Println("Num infected - ", virus.numInfected)
}

func getIpListFromFilename(filename string) [][]string {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	var gridArr [][]string
	for _, ipStr := range ipStrList {
		var nodesList []string
		for _, node := range ipStr {
			nodesList = append(nodesList, string(node))
		}
		gridArr = append(gridArr, nodesList)
	}
	return gridArr
}
