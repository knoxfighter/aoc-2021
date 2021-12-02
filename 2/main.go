package main

import (
	"bufio"
	"bytes"
	"github.com/herrnan/aoc-helper"
	"log"
	"strconv"
	"strings"
)

func main() {
	helper, err := aoc.NewHelper(aoc.Day(2), aoc.Year(2021))
	if err != nil {
		log.Fatalln(err)
	}

	input, err := helper.GetInput()
	if err != nil {
		log.Fatalln(err)
	}

	firstPuzzle(input)
	log.Println("---------------------------")
	secondPuzzle(input)
}

func firstPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	var position int
	var depth int
	for scanner.Scan() {
		text := scanner.Text()

		split := strings.Split(text, " ")
		atoi, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalln(err)
		}
		switch split[0] {
		case "forward":
			position += atoi
		case "up":
			depth -= atoi
		case "down":
			depth += atoi
		}
	}

	log.Printf("\nPosition: %d\nDepth: %d", position, depth)

	log.Printf("Result: %d", position*depth)
}

func secondPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	var position, depth, aim int
	for scanner.Scan() {
		text := scanner.Text()

		split := strings.Split(text, " ")
		atoi, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalln(err)
		}

		switch split[0] {
		case "forward":
			position += atoi
			depth += aim * atoi
		case "up":
			aim -= atoi
		case "down":
			aim += atoi
		}
	}

	log.Printf("\nPosition: %d\nDepth: %d\nAim: %d", position, depth, aim)
	log.Printf("Result: %d", position*depth)
}
