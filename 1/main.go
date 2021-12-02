package main

import (
	"bufio"
	"bytes"
	"github.com/herrnan/aoc-helper"
	"log"
	"strconv"
)

func main() {
	helper, err := aoc.NewHelper(aoc.Day(1), aoc.Year(2021))
	if err != nil {
		log.Fatalln(err)
	}

	input, err := helper.GetInput()
	if err != nil {
		log.Fatalln(err)
	}

	firstPuzzle(input)
	secondPuzzle(input)
}

func firstPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	lastDepth := 0
	var decreaseAmount int
	for scanner.Scan() {
		text := scanner.Text()
		depth, err := strconv.Atoi(text)
		if err != nil {
			log.Fatalln(err)
		}

		if lastDepth == 0 {
			// do nothing
		} else if lastDepth < depth {
			decreaseAmount++
		}

		lastDepth = depth
	}
	log.Println(decreaseAmount)
}

func secondPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)
	var file []string
	for scanner.Scan() {
		file = append(file, scanner.Text())
	}

	lastCalc := 0
	decreaseAmount := 0
	for i := 1; i < len(file)-1; i++ {
		lastLine := file[i-1]
		curLine := file[i]
		nextLine := file[i+1]

		lastLineNum, err := strconv.Atoi(lastLine)
		if err != nil {
			log.Fatalln(err)
		}

		curLineNum, err := strconv.Atoi(curLine)
		if err != nil {
			log.Fatalln(err)
		}

		nextLineNum, err := strconv.Atoi(nextLine)
		if err != nil {
			log.Fatalln(err)
		}

		totalLineNum := lastLineNum + curLineNum + nextLineNum

		if lastCalc == 0 {
			// do nothing
		} else if lastCalc < totalLineNum {
			decreaseAmount++
		}

		lastCalc = totalLineNum
	}

	log.Println(decreaseAmount)
}
