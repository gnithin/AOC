package main

import (
	"fmt"
	"strconv"
)

// Part 1
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

func getNumUsedSquares(ipStr string) (int, [][]int) {
	numRows := 128
	numUsed := 0
	used := "1"

	var rowwiseUsedIndicies [][]int

	for i := 0; i < numRows; i++ {
		var usedIndices []int
		currStr := ipStr + "-" + strconv.Itoa(i)
		knotHash := performKnotHash(currStr)

		// Convert the hex to binary
		binStr := hexToBinStr(knotHash)

		// Find num of used
		for ind, r := range binStr {
			bin := string(r)
			if bin == used {
				numUsed += 1
				usedIndices = append(usedIndices, ind)
			}
		}

		rowwiseUsedIndicies = append(rowwiseUsedIndicies, usedIndices)
	}
	return numUsed, rowwiseUsedIndicies
}

// Part 2
type Coords struct {
	row int
	col int
}

func getNumRegions(ipKey string) int {
	_, rowwiseUsedIndicies := getNumUsedSquares(ipKey)
	/*
		rowwiseUsedIndicies := [][]int{
			{0, 1, 3, 4},
			{0, 3},
			{1, 2, 3},
			{},
			{0, 4},
		}

		rowwiseUsedIndicies := [][]int{
			{0, 1, 3, 5},
			{1, 3, 5, 7},
			{4, 6},
			{0, 2, 4, 5, 7},
			{1, 2, 4},
			{0, 1, 4, 7},
			{1, 5},
			{0, 1, 3, 5, 6},
		}

	*/
	numRegions := 0
	row := 0

	for row < len(rowwiseUsedIndicies) {
		usedIndicesList := rowwiseUsedIndicies[row]

		if len(usedIndicesList) == 0 {
			row += 1
			continue
		}
		currCoords := Coords{row: row, col: usedIndicesList[0]}
		regionCoords := mapARegionForCoords(currCoords, &rowwiseUsedIndicies)

		if len(regionCoords) == 0 {
			panic("This shouldn't happen - RegionCoords cannot be empty!")
		}

		// Remove all the region coords from the rowwiseUsedIndices
		for _, regionCoord := range regionCoords {
			colsList := rowwiseUsedIndicies[regionCoord.row]
			colPos := find(regionCoord.col, colsList)
			if colPos == -1 {
				panic("This shouldn't happen - Column position cannot unmatch a region-coord")
			}
			colsList = append(colsList[:colPos], colsList[colPos+1:]...)
			rowwiseUsedIndicies[regionCoord.row] = colsList
		}
		numRegions += 1
	}
	return numRegions
}

func mapARegionForCoords(currCoords Coords, rowwiseUsedIndicies *[][]int) []Coords {
	var regionQueue []Coords
	regionQueue = append(regionQueue, currCoords)
	var regionCoords = []Coords{currCoords}

	for len(regionQueue) != 0 {
		poppedCoordsPtr := regionQueue[0]
		regionQueue = regionQueue[1:]

		// All coords
		leftCoord := Coords{
			row: poppedCoordsPtr.row,
			col: poppedCoordsPtr.col - 1,
		}
		rightCoord := Coords{
			row: poppedCoordsPtr.row,
			col: poppedCoordsPtr.col + 1,
		}
		topCoord := Coords{
			row: poppedCoordsPtr.row - 1,
			col: poppedCoordsPtr.col,
		}
		bottomCoord := Coords{
			row: poppedCoordsPtr.row + 1,
			col: poppedCoordsPtr.col,
		}
		for _, c := range []Coords{leftCoord, rightCoord, topCoord, bottomCoord} {
			// Check if c is already present in regionCoords
			isAlreadyPresent := false
			for _, regionC := range regionCoords {
				if regionC.col == c.col && regionC.row == c.row {
					isAlreadyPresent = true
					break
				}
			}
			if isAlreadyPresent {
				continue
			}

			if findCoord(c, rowwiseUsedIndicies) {
				regionCoords = append(regionCoords, c)
				regionQueue = append(regionQueue, c)
			}
		}
	}
	return regionCoords
}

func findCoord(coords Coords, rowwiseUsedIndicies *[][]int) bool {
	if coords.row < 0 || coords.col < 0 {
		return false
	}

	totalNumRows := len(*rowwiseUsedIndicies)
	if coords.row >= totalNumRows {
		return false
	}

	colsList := (*rowwiseUsedIndicies)[coords.row]
	if len(colsList) == 0 {
		return false
	}
	pos := find(coords.col, colsList)
	return pos != -1
}

func find(needle int, haystack []int) int {
	for pos, val := range haystack {
		if val == needle {
			return pos
		}
	}
	return -1
}

// Main
func main() {
	//ipKey := "flqrgnkx"
	ipKey := "jxqlasbh"

	fmt.Println("Ip key - ", ipKey)
	used, _ := getNumUsedSquares(ipKey)
	fmt.Println("P1 : Used squares -", used)

	numRegions := getNumRegions(ipKey)
	fmt.Println("P2 : Number of regions -", numRegions)
}
