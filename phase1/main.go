package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isEndCharacter(character uint8) bool {
	if character == 33 || character == 46 || character == 63 {
		return true
	} else {
		return false
	}
}

func fixCapitalization(inputLine string) string {
	var outputLine string
	outputLine = ""
	var isEnd bool
	isEnd = true

	for i := 0; i < len(inputLine); i++ {
		if isEnd {
			if inputLine[i] != ' ' {
				outputLine += strings.ToUpper(string(inputLine[i]))
				isEnd = false
			} else {
				outputLine += string(inputLine[i])
			}
		} else {
			outputLine += string(inputLine[i])
		}
		if !isEnd {
			isEnd = isEndCharacter(inputLine[i])
		}
	}

	return outputLine
}

func fixOrdinalNumbers(inputLine string) string {
	var outputLine string
	outputLine = ""

	outputLine = inputLine

	//outputLine = strings.ReplaceAll(outputLine, fmt.Sprintf("%d", 1), "1st")
	//outputLine = strings.ReplaceAll(outputLine, fmt.Sprintf("%d", 2), "2nd")
	//outputLine = strings.ReplaceAll(outputLine, fmt.Sprintf("%d", 3), "3rd")
	//for i := 4; i <= 20; i++ {
	//	//outputLine = strings.ReplaceAll(outputLine, fmt.Sprintf("%d", i), fmt.Sprintf("%dth", i))
	//	outputLine = strings.ReplaceAll(outputLine, fmt.Sprintf("%d", i), fmt.Sprintf("%dth", i))
	//}
	//outputLine = strings.ReplaceAll(outputLine, "21ST", "21ST")
	//outputLine = strings.ReplaceAll(outputLine, "22ND", "22ND")
	//outputLine = strings.ReplaceAll(outputLine, "23RD", "23RD")
	//for i := 24; i <= 30; i++ {
	//	outputLine = strings.ReplaceAll(outputLine, fmt.Sprintf("%dTH", i), fmt.Sprintf("%dTH", i))
	//}

	return outputLine
}

func fixInputLine(inputLine string) string {
	var outputLine string

	outputLine = fixCapitalization(inputLine)
	outputLine = fixOrdinalNumbers(outputLine)

	return outputLine
}

func main() {
	//var inputLine string
	var outputLine string

	//fmt.Print("Enter your input:\n ")
	//fmt.Scanf("%s", &inputLine)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	inputLine, _ := reader.ReadString('\n')

	outputLine = fixInputLine(inputLine)

	fmt.Println(outputLine)
}
