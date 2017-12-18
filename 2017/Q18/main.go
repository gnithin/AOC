package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Part 1

// Main
func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	instnList := getIpListFromFilename(filename)

	// Part 1
	interpreter := createInterpreter(instnList)
	interpreter.run()
	recoveredFreq := interpreter.recoveredFreq
	fmt.Println("Recovered freq - ", recoveredFreq)
}

func getIpListFromFilename(filename string) []Instruction {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	var instnList []Instruction
	for _, ipStr := range ipStrList {
		ipStr = strings.TrimSpace(ipStr)
		ipComponents := strings.Split(ipStr, " ")
		cmdName := ipComponents[0]

		register := ""
		argIndex := 1
		if len(ipComponents) > 2 {
			register = ipComponents[1]
			argIndex = 2
		}
		var argument interface{}
		value, err := strconv.Atoi(ipComponents[argIndex])
		if err != nil {
			argument = ipComponents[argIndex]
		} else {
			argument = value
		}

		instn := Instruction{
			cmdName,
			register,
			argument,
		}
		instnList = append(instnList, instn)
	}
	return instnList
}
