package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	// Read input from file
	filename := "ip.txt"
	contentsByteList, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	contentsStr := string(contentsByteList)
	contentsStr = strings.TrimSpace(contentsStr)

	ipList := strings.Split(contentsStr, "\n")
	fmt.Println(len(ipList))

	validCount := 0
	for _, ip := range ipList {
		isValid := isPassphraseValid(ip)
		fmt.Println(isValid, " - ", ip)
		if isValid {
			validCount += 1
		}
	}

	fmt.Println("Number of valid passphrases - ", validCount)
}

func isPassphraseValid(passphrase string) bool {
	phraseComponents := strings.Fields(passphrase)
	phraseMap := make(map[string]bool)
	sortedPhrasesList := []string{}

	// This for part-2
	for _, pc := range phraseComponents {
		s := strings.Split(pc, "")
		sort.Strings(s)
		sortedStr := strings.Join(s, "")
		sortedPhrasesList = append(sortedPhrasesList, sortedStr)
	}

	for _, pc := range sortedPhrasesList {
		isFound, _ := phraseMap[pc]
		if isFound {
			return false
		}
		phraseMap[pc] = true
	}
	return true
}
