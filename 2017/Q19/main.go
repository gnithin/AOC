package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

const (
	DIRN_ANCHOR = "+"
	HOR_PATH    = "-"
	VER_PATH    = "|"
	DIRN_LEFT   = "l"
	DIRN_RIGHT  = "r"
	DIRN_UP     = "u"
	DIRN_DOWN   = "d"
)

// Part 1
func getLetters(pathList []string) string {
	lettersFound := ""

	// Find the initial position
	currRow := 0
	currCol := findPos(VER_PATH, pathList[currRow])
	if currCol == -1 {
		panic("Could not find initial position")
	}

	// Traverse the path
	currDirn := DIRN_DOWN
	endFound := false

	for {
		currPath := string(pathList[currRow][currCol])
		//fmt.Println("Curr path - ", currPath)

		switch currPath {
		case VER_PATH:
			// pass
		case HOR_PATH:
			// pass
		case DIRN_ANCHOR:
			//fmt.Println("Dirn change - ")
			//fmt.Println("Old dirn - ", currDirn)
			// Find out where to go based on the current direction
			if currDirn == DIRN_UP || currDirn == DIRN_DOWN {
				// Examine left and right
				leftCol := currCol - 1
				if leftCol >= 0 &&
					(string(pathList[currRow][leftCol]) == HOR_PATH ||
						unicode.IsLetter(rune(pathList[currRow][leftCol]))) {
					currDirn = DIRN_LEFT
				} else {
					currDirn = DIRN_RIGHT
				}
			} else {
				// Examine top and bottom
				topRow := currRow - 1
				if topRow >= 0 &&
					(string(pathList[topRow][currCol]) == VER_PATH ||
						unicode.IsLetter(rune(pathList[topRow][currCol]))) {
					currDirn = DIRN_UP
				} else {
					currDirn = DIRN_DOWN
				}
			}
			//fmt.Println("new dirn - ", currDirn)
		default:
			if unicode.IsLetter([]rune(currPath)[0]) {
				lettersFound += currPath
			} else {
				endFound = true
			}
		}

		if endFound {
			break
		}

		// Update the path count
		switch currDirn {
		case DIRN_UP:
			currRow -= 1
		case DIRN_DOWN:
			currRow += 1
		case DIRN_LEFT:
			currCol -= 1
		case DIRN_RIGHT:
			currCol += 1
		}
	}

	return lettersFound
}

func findPos(needle string, haystack string) int {
	for pos, r := range haystack {
		if needle == string(r) {
			return pos
		}
	}
	return -1
}

// Main
func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	ipList := getIpListFromFilename(filename)
	/*
		for _, ip := range ipList {
			fmt.Println(ip)
		}
	*/
	letters := getLetters(ipList)
	fmt.Println("Letters from path - ", letters)
}

func getIpListFromFilename(filename string) []string {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	ipStrList := strings.Split(fileContents, "\n")
	return ipStrList
}
