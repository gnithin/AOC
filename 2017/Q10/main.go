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
) ([]int, int, int) {
	ipArr := make([]int, len(ipArrOriginal))
	copy(ipArr, ipArrOriginal)

	for _, currLen := range lenList {
		reverseSlice(ipArr, currPos, currLen)
		currPos = (currPos + currLen + skip) % len(ipArr)
		skip += 1
	}

	return ipArr, currPos, skip
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

// Part 2
func performKnotHash(ipArr, lenList []int) string {
	pos := 0
	skip := 0
	numRounds := 64

	for i := 0; i < numRounds; i++ {
		ipArr, pos, skip = performTransformations(ipArr, lenList, pos, skip)
	}
	return convertIntListToHex(createDenseHashFromSparseHash(ipArr))
}

func createDenseHashFromSparseHash(sparseHash []int) []int {
	divisor := 16
	if len(sparseHash)%divisor != 0 {
		panic("Invalid sparse hash!")
	}

	var denseHash []int

	for i := 0; i < len(sparseHash); i += divisor {
		selectedSlice := sparseHash[i : i+divisor]

		// Find the xor of slice
		finalXor := selectedSlice[0]
		for j := 1; j < len(selectedSlice); j++ {
			finalXor = finalXor ^ selectedSlice[j]
		}
		denseHash = append(denseHash, finalXor)
	}

	return denseHash
}

func convertIntListToHex(intList []int) string {
	var hexStringList []string
	for _, ip := range intList {
		ipHex := strconv.FormatInt(int64(ip), 16)
		if len(ipHex) == 1 {
			ipHex = "0" + ipHex
		}
		hexStringList = append(hexStringList, ipHex)
	}

	return strings.Join(hexStringList, "")
}

// Main
func main() {
	filename := "ip.txt"
	ipArrLen := 256

	part1(filename, ipArrLen)
	part2(filename, ipArrLen)
}

// Part 2
func part2(filename string, ipArrLen int) {
	lenList := getFormattedIpListFromFilename(filename)
	var ipArr []int
	for i := 0; i < ipArrLen; i++ {
		ipArr = append(ipArr, i)
	}

	hashStr := performKnotHash(ipArr, lenList)
	fmt.Println("Final knot hash -", hashStr)
}

func getFormattedIpListFromFilename(filename string) []int {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)

	var ipIntList []int
	for _, c := range fileContents {
		ipIntList = append(ipIntList, int(c))
	}

	ipIntList = append(ipIntList, []int{17, 31, 73, 47, 23}...)
	return ipIntList
}

// Part 1
func part1(filename string, ipArrLen int) {
	lenList := getIpListFromFilename(filename)
	var ipArr []int
	for i := 0; i < ipArrLen; i++ {
		ipArr = append(ipArr, i)
	}
	startingPos := 0
	initialSkip := 0

	transformedArr, _, _ := performTransformations(ipArr, lenList, startingPos, initialSkip)
	fmt.Println("P1 Final Product - ", transformedArr[0]*transformedArr[1])
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
