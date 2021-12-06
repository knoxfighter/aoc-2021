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
	helper, err := aoc.NewHelper(aoc.Day(6), aoc.Year(2021))
	if err != nil {
		log.Fatalln(err)
	}

	input, err := helper.GetInput()
	if err != nil {
		log.Fatalln(err)
	}

	firstPuzzle(input)
}

func firstPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	var lanternfish [9]int

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")

		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalln(err)
			}

			lanternfish[i]++
		}
	}

	for i := 0; i < 256; i++ {
		fish := lanternfish[0]
		lanternfish[0] = lanternfish[1]
		lanternfish[1] = lanternfish[2]
		lanternfish[2] = lanternfish[3]
		lanternfish[3] = lanternfish[4]
		lanternfish[4] = lanternfish[5]
		lanternfish[5] = lanternfish[6]
		lanternfish[6] = lanternfish[7] + fish
		lanternfish[7] = lanternfish[8]
		lanternfish[8] = fish

		if i == 79 {
			var amount int
			for _, i := range lanternfish {
				amount += i
			}
			log.Printf("after 80 days: %d", amount)
		}
	}

	var amount int
	for _, i := range lanternfish {
		amount += i
	}
	log.Printf("after 256 days: %d", amount)
}
