package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
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
			isEnd = isEndCharacter(inputLine[i])
		}
	}

	return outputLine
}

func ordinal(n int) string {
	var suffix string
	suffix = ""

	if 11 <= (n%100) && (n%100) <= 13 {
		suffix = "th"
	} else {
		a := [5]string{"th", "st", "nd", "rd", "th"}
		suffix = a[int(math.Min(float64(int((n%10))), 4))]
	}

	return suffix
}

func fixOrdinalNumbers(inputLine string) string {
	var outputLine string
	outputLine = ""
	var digit string
	digit = ""

	for i := 0; i < len(inputLine); i++ {
		if unicode.IsDigit(rune(inputLine[i])) {
			digit += string(inputLine[i])
		}

		outputLine += string(inputLine[i])

		if digit != "" && !unicode.IsDigit(rune(inputLine[i+1])) {
			mark, _ := strconv.Atoi(digit)
			outputLine += ordinal(mark)
			digit = ""
		}
	}

	return outputLine
}

func fixInputLine(inputLine string) string {
	var outputLine string

	outputLine = fixCapitalization(inputLine)
	outputLine = fixOrdinalNumbers(outputLine)

	return outputLine
}

func main() {
	var outputLine string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	inputLine, _ := reader.ReadString('\n')

	outputLine = fixInputLine(inputLine)

	fmt.Println(outputLine)
}
