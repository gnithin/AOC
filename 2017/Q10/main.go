package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Part 1
func performTransformations(
	ipArrOriginal, lenList []int,
	currPos, skip int,
) []int {
	ipArr := make([]int, len(ipArrOriginal))
	copy(ipArr, ipArrOriginal)

	for _, currLen := range lenList {
		//fmt.Println("*******")
		//fmt.Println("Pos ", currPos, " len ", currLen)
		//fmt.Println("Original ", ipArr)
		reverseSlice(ipArr, currPos, currLen)
		//fmt.Println("Reversed ", ipArr)
		currPos = (currPos + currLen + skip) % len(ipArr)
		skip += 1
	}

	return ipArr
}

func reverseSlice(ipSlice []int, currPos, currLen int) {
	var selectedSlice []int
	if currPos+currLen > len(ipSlice) {
		sliceLenFromStart := currLen - (len(ipSlice) - currPos)
		selectedSlice = append(ipSlice[currPos:], ipSlice[:sliceLenFromStart]...)
	} else {
		selectedSlice = ipSlice[currPos:(currPos + currLen)]
	}
	// Reverse the slice
	lenSlice := len(selectedSlice)
	for i, j := 0, lenSlice-1; i < lenSlice/2; i, j = i+1, j-1 {
		selectedSlice[i], selectedSlice[j] = selectedSlice[j], selectedSlice[i]
	}

	pos := currPos
	for _, val := range selectedSlice {
		ipSlice[pos] = val
		pos = (pos + 1) % len(ipSlice)
	}
}

// Main
func main() {
	filename := "ip.txt"
	ipArrLen := 256

	//filename := "trial.txt"
	//ipArrLen := 5

	lenList := getIpListFromFilename(filename)
	var ipArr []int
	for i := 0; i < ipArrLen; i++ {
		ipArr = append(ipArr, i)
	}
	startingPos := 0
	initialSkip := 0

	transformedArr := performTransformations(ipArr, lenList, startingPos, initialSkip)
	//fmt.Println("Final transformed - ", transformedArr)
	fmt.Println("Final Product - ", transformedArr[0]*transformedArr[1])
}

func getIpListFromFilename(filename string) []int {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, ",")
	var ipIntList []int
	for _, ip := range ipStrList {
		val, _ := strconv.Atoi(strings.TrimSpace(string(ip)))
		ipIntList = append(ipIntList, val)
	}
	return ipIntList
}
