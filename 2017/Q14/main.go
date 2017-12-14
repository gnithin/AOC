package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func __performKnotHash(ipArr, lenList []int) string {
	pos := 0
	skip := 0
	numRounds := 64

	for i := 0; i < numRounds; i++ {
		ipArr, pos, skip = performTransformations(ipArr, lenList, pos, skip)
	}
	return convertIntListToHex(createDenseHashFromSparseHash(ipArr))
}

func performKnotHash(ipStr string) string {
	ipLen := 256
	var lenList []int
	for i := 0; i < ipLen; i++ {
		lenList = append(lenList, i)
	}

	// Convert the ipstr to ascii values
	var ipArr = []int{}
	for _, r := range ipStr {
		ipArr = append(ipArr, int(r))
	}
	ipArr = append(ipArr, []int{17, 31, 73, 47, 23}...)
	return __performKnotHash(lenList, ipArr)
}

// Part 1
func getNumUsedSquares(ipStr string) int {
	numRows := 128
	numUsed := 0
	used := "1"

	for i := 0; i < numRows; i++ {
		currStr := ipStr + "-" + strconv.Itoa(i)
		knotHash := performKnotHash(currStr)

		// Convert the hex to binary
		binStr := hexToBinStr(knotHash)

		// Find num of used
		for _, r := range binStr {
			bin := string(r)
			if bin == used {
				numUsed += 1
			}
		}
	}
	return numUsed
}

func hexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}

	// %016b indicates base 2, zero padded, with 16 characters
	return fmt.Sprintf("%04b", ui), nil
}

func hexToBinStr(hex string) string {
	binStr := ""
	for _, r := range hex {
		currBinStr, e := hexToBin(string(r))
		if e != nil {
			panic(e)
		}
		binStr += currBinStr
	}
	return binStr
}

// Main
func main() {
	//ipKey := "flqrgnkx"
	ipKey := "jxqlasbh"
	fmt.Println("Ip key - ", ipKey)
	used := getNumUsedSquares(ipKey)
	fmt.Println("Used squares - ", used)
}
