package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Part 1
func findNumProgramsConnectedTo(idStr string, ipMap map[string][]string) (int, []string) {
	var visited []string
	var queue []string

	queue = append(queue, idStr)
	for len(queue) > 0 {
		key := queue[0]
		queue = queue[1:]

		if isItemInArray(key, visited) {
			continue
		}
		visited = append(visited, key)

		destList, ok := ipMap[key]
		if !ok || len(destList) == 0 {
			continue
		}

		for _, element := range destList {
			if isItemInArray(element, visited) {
				continue
			}
			queue = append(queue, element)
		}
	}

	return len(visited), visited
}

func isItemInArray(needle string, haystack []string) bool {
	return indexOfItemInArray(needle, haystack) != -1
}

func indexOfItemInArray(needle string, haystack []string) int {
	for pos, item := range haystack {
		if item == needle {
			return pos
		}
	}
	return -1
}

// Part 2
func findNumGroups(ipMap map[string][]string) int {
	numGroups := 0
	var keysList []string
	for key, _ := range ipMap {
		keysList = append(keysList, key)
	}

	for len(keysList) > 0 {
		item := keysList[0]
		_, groupMembers := findNumProgramsConnectedTo(item, ipMap)
		keysList = updateSliceByRemovingItems(keysList, groupMembers)
		numGroups += 1
	}
	return numGroups
}

func updateSliceByRemovingItems(slice, removalItems []string) []string {
	for _, item := range removalItems {
		pos := indexOfItemInArray(item, slice)
		if pos == -1 {
			continue
		}
		slice = append(slice[:pos], slice[pos+1:]...)
	}
	return slice
}

// Main
func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	ipMap := getIpListFromFilename(filename)

	idStr := "0"
	numProgs, _ := findNumProgramsConnectedTo(idStr, ipMap)
	fmt.Println("Number of programs connected to", idStr, "-", numProgs)

	numGroups := findNumGroups(ipMap)
	fmt.Println("Number of groups - ", numGroups)
}

func getIpListFromFilename(filename string) map[string][]string {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	fileContents = strings.Replace(fileContents, " ", "", -1)
	ipStrList := strings.Split(fileContents, "\n")

	ipMap := make(map[string][]string)
	parentSep := "<->"
	childSep := ","
	for _, ip := range ipStrList {
		keyValuesList := strings.Split(ip, parentSep)
		key := keyValuesList[0]
		valueStr := keyValuesList[1]
		valuesList := strings.Split(valueStr, childSep)
		ipMap[key] = valuesList
	}

	return ipMap
}
