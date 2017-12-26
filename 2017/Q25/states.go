package main

import (
	"fmt"
	re "regexp"
	"strconv"
	"strings"
)

type cursorDirn int

const (
	L = iota
	R
)

type State struct {
	stateName     string
	predicateList []Predictate
}

type Predictate struct {
	predicateCurrVal int
	writeVal         int
	stepSize         int
	dirn             cursorDirn
	nextState        string
}

func CreateStateFromIpStr(ipStr string) State {
	currState := State{}
	lines := strings.Split(ipStr, "\n")
	// Parse the current state
	stateNameRegex := re.MustCompile(`state ([A-Z])`)
	matches := stateNameRegex.FindStringSubmatch(string(lines[0]))
	currState.stateName = matches[1]
	predicateLines := lines[1:]

	var predicateList []Predictate
	for i := 0; i < len(predicateLines); i = i + 4 {
		predicateListStr := predicateLines[i : i+4]
		predicate := Predictate{}
		rawText := strings.Join(predicateListStr, "\n")

		predicateRegex := re.MustCompile(`If .*?(\d+):\n.*Write.*?(\d+)\.\n.*?Move\s+(\w+).*?the\s+(\w+)\.\n.*-\s+.*?state\s+([A-Z]+)\.`)

		matches := predicateRegex.FindStringSubmatch(rawText)
		predicate.predicateCurrVal, _ = strconv.Atoi(matches[1])
		predicate.writeVal, _ = strconv.Atoi(matches[2])
		predicate.stepSize = 1
		if matches[4] == "left" {
			predicate.dirn = L
		} else {
			predicate.dirn = R
		}
		predicate.nextState = matches[5]
		predicateList = append(predicateList, predicate)
	}
	currState.predicateList = predicateList

	fmt.Printf("")
	return currState
}

type StateManager struct {
	diagnosticPeriod int
	initialState     string
	stateDetails     map[string]State
}

func createStateManager(
	initialState string,
	numDiagRuns int,
	statesList []State,
) StateManager {
	// Create a map from the states list
	stateDetails := make(map[string]State)
	for _, state := range statesList {
		stateDetails[state.stateName] = state
	}
	return StateManager{
		diagnosticPeriod: numDiagRuns,
		initialState:     initialState,
		stateDetails:     stateDetails,
	}
}

type TuringMachine struct {
	currPosition  int
	stateMgr      StateManager
	oneValuesMap  map[int]int
	currStateName string
}

func CreateTuringMachine(stateMgr StateManager) TuringMachine {
	turningMachine := TuringMachine{
		stateMgr:      stateMgr,
		currStateName: stateMgr.initialState,
		currPosition:  0,
		oneValuesMap:  make(map[int]int),
	}
	return turningMachine
}

func (self *TuringMachine) runForOneCycle() {
	for i := 0; i < self.stateMgr.diagnosticPeriod; i++ {
		/*
			fmt.Println("*********")
			fmt.Println("Run - ", i+1)
			fmt.Println("Curr pos - ", self.currPosition)
			fmt.Println("Curr statename - ", self.currStateName)
			fmt.Println("Num ones - ", len(self.oneValuesMap))
		*/
		self.run()
	}
}

func (self *TuringMachine) run() {
	currPos := self.currPosition
	currVal := 0
	_, oneFound := self.oneValuesMap[currPos]
	if oneFound {
		currVal = 1
	}

	currState := self.currStateName
	stateDetails := self.stateMgr.stateDetails[currState]
	predicateList := stateDetails.predicateList
	predicateFound := false
	for _, p := range predicateList {
		if p.predicateCurrVal != currVal {
			continue
		}
		predicateFound = true
		writeVal := p.writeVal
		if writeVal == 1 {
			self.oneValuesMap[currPos] = 1
		} else {
			delete(self.oneValuesMap, currPos)
		}
		if p.dirn == L {
			self.currPosition -= 1
		} else {
			self.currPosition += 1
		}
		self.currStateName = p.nextState
		break
	}

	if predicateFound != true {
		panic("Cannot find predicate!")
	}
}
