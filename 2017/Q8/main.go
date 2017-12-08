package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var REGISTER_STR string = "iRegister"
var OPCODE_STR string = "iOpCode"
var OPVAL_STR string = "iOpVal"
var CONDN_FULL_STR string = "iCondnFull"
var CONDN_REGISTER_STR string = "iCondnRegister"
var CONDN_COMPARATOR_STR string = "iCondnComparator"
var CONDN_VAL_STR string = "iCondnVal"
var INC string = "inc"
var DEC string = "dec"
var GLOBAL_MAX_VAL = -1

var GLOBAL_VAR_MAP map[string]int = make(map[string]int)

type Condition struct {
	register   string
	comparator string
	val        int
}

type Instruction struct {
	register  string
	opCode    string
	opVal     int
	condition Condition
}

func (instn *Instruction) getInstructionFromRawStr(rawInstructionStr string) {
	// regexStr := `^(\w+)\s+(inc|dec)\s+([0-9-]+)\s+if\s+((\w+)\s+([!><=]+)\s+([0-9-]+))`
	// Regex details - https://regex101.com/r/IPrhl4/1
	regexStr := `^(?P<iRegister>\w+)\s+(?P<iOpCode>inc|dec)\s+(?P<iOpVal>[0-9-]+)\s+if\s+(?P<iCondnFull>(?P<iCondnRegister>\w+)\s+(?P<iCondnComparator>[!><=]+)\s+(?P<iCondnVal>[0-9-]+))`
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		panic(err)
	}

	allMatchesList := regex.FindAllStringSubmatch(rawInstructionStr, -1)
	if len(allMatchesList) < 1 {
		panic("Regex match failed!")
	}
	matchesList := allMatchesList[0]

	subExpNames := regex.SubexpNames()
	matchesMap := map[string]string{}
	//fmt.Println("***")
	for i, match := range matchesList {
		subExpName := subExpNames[i]
		if subExpName == "" {
			continue
		}
		matchesMap[subExpName] = match
		//fmt.Println(subExpName, ": ", match)
	}

	instn.register = matchesMap[REGISTER_STR]
	instn.opCode = matchesMap[OPCODE_STR]
	instn.opVal, _ = strconv.Atoi(matchesMap[OPVAL_STR])

	condition := Condition{}
	condition.register = matchesMap[CONDN_REGISTER_STR]
	condition.comparator = matchesMap[CONDN_COMPARATOR_STR]
	condition.val, _ = strconv.Atoi(matchesMap[CONDN_VAL_STR])
	instn.condition = condition
}

func (instn *Instruction) evaluate() {
	// Evaluate the condition first
	cRegName := instn.condition.register
	cRegVal, _ := GLOBAL_VAR_MAP[cRegName]

	cComparator := instn.condition.comparator
	cExpectedVal := instn.condition.val
	conditionResult := false

	switch cComparator {
	case ">":
		conditionResult = cRegVal > cExpectedVal
	case "<":
		conditionResult = cRegVal < cExpectedVal
	case "<=":
		conditionResult = cRegVal <= cExpectedVal
	case ">=":
		conditionResult = cRegVal >= cExpectedVal
	case "==":
		conditionResult = cRegVal == cExpectedVal
	case "!=":
		conditionResult = cRegVal != cExpectedVal
	}

	if conditionResult {
		regVal := GLOBAL_VAR_MAP[instn.register]
		opVal := instn.opVal
		if instn.opCode == INC {
			regVal += opVal
		} else {
			regVal -= opVal
		}
		GLOBAL_VAR_MAP[instn.register] = regVal

		if regVal > GLOBAL_MAX_VAL {
			GLOBAL_MAX_VAL = regVal
		}
	}
}

var instructionsList []*Instruction

// Part 1
func parseIpList(ipList []string) {
	for _, ip := range ipList {
		rawInstructionStr := strings.TrimSpace(ip)
		instruction := Instruction{}
		instruction.getInstructionFromRawStr(rawInstructionStr)
		instructionsList = append(instructionsList, &instruction)
	}
}

func evaluateInstructions() {
	for _, instruction := range instructionsList {
		instruction.evaluate()
	}
}

func getMaxValueOfVariables() int {
	maxVal := 0
	for _, val := range GLOBAL_VAR_MAP {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// Main

func main() {
	filename := "ip.txt"
	ipList := getIpListFromFilename(filename)
	parseIpList(ipList)
	evaluateInstructions()
	maxVal := getMaxValueOfVariables()
	fmt.Println("P1: Max val - ", maxVal)
	fmt.Println("P2: Max value held all the time - ", GLOBAL_MAX_VAL)
}

func getIpListFromFilename(filename string) []string {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	return ipStrList
}
