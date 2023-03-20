package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
	"sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

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

type job struct {
    line string
    lineNumber int
}

type jobRes struct {
    repairedLine string
    lineNumber int
}

func doJob(inCh chan job, outCh chan jobRes) {
	for toDoJob := range inCh {
		outCh <- jobRes{repairedLine: fixInputLine(toDoJob.line), lineNumber: toDoJob.lineNumber}
	}
}

func main() {
	// get number of workers

	var n int
	fmt.Print("Enter a number of workers: ")
	fmt.Scanf("%d", &n)

	// initialize routines and channels
	inChannel := make(chan job, 100)
	outChannel := make(chan jobRes, 100)

	for i := 0; i < n; i++ {
		go doJob(inChannel, outChannel);
	}

	// open file and read from it
	f, err := os.Open("./input.txt")
    check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var busyWorkersCount int = 0

	var results []jobRes;
	var i int = 0
	var j int = 0

	for scanner.Scan() {
		inChannel <- job{line: scanner.Text(), lineNumber: j}
		j += 1
		busyWorkersCount += 1

		if busyWorkersCount >= n {
			tmpJob := <- outChannel
			results = append(results, tmpJob)
			i += 1
		}
    }
	
	for i < j {
		tmpJob := <- outChannel
		results = append(results, tmpJob)
		i += 1
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].lineNumber < results[j].lineNumber
	})

	//
	w, err := os.OpenFile("./output.txt", os.O_CREATE | os.O_WRONLY, 0644)

	check(err)

	datawriter := bufio.NewWriter(w)

	for _, jobRes := range results {
		datawriter.WriteString(jobRes.repairedLine + "\n")
	}
	datawriter.Flush()
}
