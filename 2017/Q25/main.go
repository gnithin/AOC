package main

import (
	"fmt"
	"io/ioutil"
	re "regexp"
	"strconv"
	"strings"
)

func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	//ipList := getIpListFromFilename(filename)
	//fmt.Printf("Input - [%s]\n", strings.Join(ipList, ", "))
	stateManager := getStateManagerFromFile(filename)
	fmt.Println(stateManager.diagnosticPeriod)
	fmt.Println(stateManager.initialState)
	fmt.Println("*******")
	/*
		for k, v := range stateManager.stateDetails {
			fmt.Println("####")
			fmt.Println(k, " => ", v)
			fmt.Println("####")
		}
	*/

	// Run for one diagnostic period
	turningMachine := CreateTuringMachine(stateManager)
	turningMachine.runForOneCycle()
	fmt.Println("*******")
	fmt.Println("Num ones - ", len(turningMachine.oneValuesMap))
}

func getStateManagerFromFile(filename string) StateManager {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n\n")

	beginBlock := strings.Split(ipStrList[0], "\n")
	// Get the begin state -
	beginRe := `state ([A-Z])`
	beginRegex := re.MustCompile(beginRe)
	beginMatches := beginRegex.FindStringSubmatch(beginBlock[0])
	beginState := beginMatches[1]
	if beginState == "" {
		panic("Cannot get begin state")
	}

	// Get diagnostic checksum
	diagRegex := re.MustCompile(`(\d+) steps\.`)
	diagMatches := diagRegex.FindStringSubmatch(beginBlock[1])
	numDiagRunsStr := diagMatches[1]

	numDiagRuns, diagErr := strconv.Atoi(numDiagRunsStr)
	if diagErr != nil || numDiagRuns == 0 {
		panic("Cannot get number of diagnostic runs")
	}
	var statesList []State
	for _, ipStr := range ipStrList[1:] {
		state := CreateStateFromIpStr(ipStr)
		statesList = append(statesList, state)
	}
	stateManager := createStateManager(beginState, numDiagRuns, statesList)
	return stateManager
}
