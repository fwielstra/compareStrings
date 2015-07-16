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
		if strings.Contains(line, ";;") {
			fmt.Println("ERROR: Line contains double semicolons!", line)
			os.Exit(1)
		}
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

	totalLeft := 0

	for key, _ := range left {
		rightValue := right[key]
		if rightValue == "" {
			totalLeft++
			fmt.Println("Key", key, "defined in", leftFile, "not found in", rightFile)
		}
	}

	totalRight := 0
	for key, _ := range right {
		leftValue := left[key]
		if leftValue == "" {
			totalRight++
			fmt.Println("Key", key, "defined in", rightFile, "not found in", leftFile)
		}
	}
	fmt.Println()
	fmt.Println(totalLeft, "keys defined in", leftFile, "not found in", rightFile)
	fmt.Println(totalRight, "keys defined in", rightFile, "not found in", leftFile)
	fmt.Println("Total difference: ", totalLeft+totalRight)

}
