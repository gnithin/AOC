package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Read the input from the file
	filename := "ip.txt"
	ok, fileContents := readContentsFromFilename(filename)
	if ok == false {
		panic(fmt.Sprintln("Couldn't read file - ", filename))
	}
	fileContents = strings.TrimSpace(fileContents)

	repetitiveDigitsList := findRepetitiveDigits(fileContents)

	sumVal := 0
	for _, s := range repetitiveDigitsList {
		sVal, _ := strconv.Atoi(string(s))
		sumVal += sVal
	}
	fmt.Println("Sum is -", sumVal)
}

func readContentsFromFilename(filename string) (bool, string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file - ", err)
		return false, ""
	}

	return true, string(data)
}

func findRepetitiveDigits(ipStr string) []string {
	repetitiveDigitsList := []string{}
	sLen := len(ipStr)

	for i := 0; i < sLen; i++ {
		currVal := ipStr[i]
		nextVal := ipStr[(i+1)%sLen]
		if currVal == nextVal {
			repetitiveDigitsList = append(repetitiveDigitsList, string(currVal))
		}
	}
	return repetitiveDigitsList
}
