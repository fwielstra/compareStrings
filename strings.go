package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// usage: compareStrings en.lproj/Localizable.strings nl.lproj/Localizable.strings

func getLines(filePath string) map[string]string {
	file, _ := os.Open(filePath)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " = ")
		if len(split) == 2 {
			result[split[0]] = split[1]
		}
	}

	return result
}

func main() {
	leftFile := os.Args[1]
	rightFile := os.Args[2]

	fmt.Println("Comparing", leftFile, "with", rightFile)
	left := getLines(leftFile)
	right := getLines(rightFile)

	total := 0

	for key, _ := range left {
		rightValue := right[key]
		if rightValue == "" {
			total++
			fmt.Println("Key", key, "defined in", os.Args[1], "not found in", os.Args[2])
		}
	}

	for key, _ := range right {
		leftValue := left[key]
		if leftValue == "" {
			total++
			fmt.Println("Key", key, "defined in", os.Args[2], "not found in", os.Args[1])
		}
	}

	fmt.Println("Total difference: ", total)

}
