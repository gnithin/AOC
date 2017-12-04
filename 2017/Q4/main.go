package main

import (
	"fmt"
	"io/ioutil"
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

	isValid := true
	for _, pc := range phraseComponents {
		isFound, _ := phraseMap[pc]
		if isFound {
			isValid = false
			break
		}
		phraseMap[pc] = true
	}
	return isValid
}
