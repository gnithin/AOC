package main

import (
	"fmt"
)

// Component
type Component struct {
	port1 int
	port2 int
}

func CreateComponentWithPorts(portA, portB int) Component {
	port1 := portA
	port2 := portB

	if port1 > port2 {
		port1, port2 = port2, port1
	}
	fmt.Printf("")
	return Component{
		port1: port1,
		port2: port2,
	}
}

// Part 1 solution
func getMaxScoreFromCompList(compList []Component) int {
	initialParentPort := 0
	initialScore := 0
	maxScore := getMaxScore(compList, initialParentPort, initialScore)
	return maxScore
}

func getMaxScore(remainingItems []Component, parentPort int, score int) int {
	// Find the components whose parentPort matches
	parentPortsIndices := removePortsFromList(parentPort, remainingItems)
	if len(parentPortsIndices) == 0 {
		return score
	}

	tempMaxScore := score
	for _, ppIndex := range parentPortsIndices {
		compWithParentPort := remainingItems[ppIndex]
		var compWithoutParentPorts []Component
		for i, v := range remainingItems {
			if i != ppIndex {
				compWithoutParentPorts = append(compWithoutParentPorts, v)
			}
		}

		otherPortVal := compWithParentPort.port1
		if otherPortVal == parentPort {
			otherPortVal = compWithParentPort.port2
		}
		currScore := score + parentPort + otherPortVal
		updatedScore := getMaxScore(compWithoutParentPorts, otherPortVal, currScore)
		if updatedScore > tempMaxScore {
			tempMaxScore = updatedScore
		}
	}

	return tempMaxScore
}

func removePortsFromList(port int, origCompList []Component) []int {
	//fmt.Println("Remove  - ", port, " -> ", origCompList)
	var foundPortsIndex []int

	for pos, comp := range origCompList {
		if comp.port1 == port || comp.port2 == port {
			foundPortsIndex = append(foundPortsIndex, pos)
		}
	}
	//fmt.Println("Founds ports index - ", foundPortsIndex)
	return foundPortsIndex
}
