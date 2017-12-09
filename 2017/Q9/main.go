package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const OPEN_GROUP = "{"
const CLOSE_GROUP = "}"
const OPEN_GARBAGE = "<"
const CLOSE_GARBAGE = ">"

// Part 1
func findScoreForGroupStr(ip string) (int, int) {
	// NOTE: In hind-sight, the idea to remove garbage like this was really stupid.
	// Should've used stack based approach. I really got behind the regex way of working :(
	// Anyways, writing this makes me appreciate regexes a lot more, especially negative look-behinds :)

	// Remove garbage
	cleanedIp, numGarbageChars := removeGarbage(ip)

	finalScore := getScoreForStr(cleanedIp)
	return finalScore, numGarbageChars
}

func removeGarbage(ip string) (string, int) {
	var modifiedStrList []rune
	numGarbageChars := 0
	i := 0
	for i < len(ip) {
		if string(ip[i]) != OPEN_GARBAGE {
			modifiedStrList = append(modifiedStrList, rune(ip[i]))
			i += 1
		} else {
			j := i + 1
			for j < len(ip) {
				if string(ip[j]) == CLOSE_GARBAGE {
					if string(ip[j-1]) != "!" {
						break
					} else {
						// Find the number of !
						negCount := 0
						for k := j - 1; k > 0; k -= 1 {
							if string(ip[k]) != "!" {
								break
							}
							negCount += 1
						}
						if negCount%2 == 0 {
							break
						}
					}
				}
				j += 1
			}
			numGarbageChars += countGarbage(ip[i+1 : j])
			i = j + 1
		}
	}

	return string(modifiedStrList), numGarbageChars
}

func countGarbage(garbage string) int {
	count := 0
	i := 0
	for i < len(garbage) {
		if string(garbage[i]) == "!" {
			i = i + 2
		} else {
			count += 1
			i = i + 1
		}
	}
	return count
}

var tagStack []string

func getScoreForStr(ip string) int {
	// Find groups
	score := 0
	for i := 0; i < len(ip); i += 1 {
		currVal := string(ip[i])
		if currVal == OPEN_GROUP {
			tagStack = append(tagStack, OPEN_GROUP)
		} else if currVal == CLOSE_GROUP {
			score += len(tagStack)
			tagStack = tagStack[:len(tagStack)-1]
		}
	}

	return score
}

// Main
func main() {
	filename := "ip.txt"
	//filename := "old_ip.txt"
	ipList := getIpListFromFilename(filename)
	for _, ip := range ipList {
		//fmt.Println("****")
		//fmt.Println(ip)
		score, numGarbageChars := findScoreForGroupStr(ip)
		fmt.Println("Score - ", score)
		fmt.Println("Garbage chars - ", numGarbageChars)
	}
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
