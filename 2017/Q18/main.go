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
	//filename := "p2.txt"
	instnList := getIpListFromFilename(filename)

	/*
		For part-1, rename the executor.go.bak to executor.go
		and uncomment this part, and rename interpreter.go to
		interpreter.go.bak and comment out part-2
	*/
	/*
		// Part 1
		interpreter := createInterpreter(instnList)
		interpreter.run()
		recoveredFreq := interpreter.recoveredFreq
		fmt.Println("Recovered freq - ", recoveredFreq)
	*/

	// Part 2
	size := 10000000
	ch1 := make(chan int, size)
	ch2 := make(chan int, size)
	responseChan := make(chan int)
	/*
		flagValue1 := false
		flagValue2 := false
	*/

	interpreter0 := createInterpreter(
		0,
		instnList,
		ch1,
		ch2,
		responseChan,
	)

	interpreter1 := createInterpreter(
		1,
		instnList,
		ch2,
		ch1,
		responseChan,
	)

	go func() {
		interpreter0.run()
	}()
	go func() {
		interpreter1.run()
	}()

	progId := <-responseChan
	if progId == 0 {
		interpreter1.stopInterrupt = true
	} else {
		interpreter0.stopInterrupt = true
	}
	_ = <-responseChan
	fmt.Println("Send value of interpreter 1 -", interpreter1.sendCount)
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
