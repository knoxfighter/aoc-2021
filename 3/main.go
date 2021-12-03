package main

import (
	"bufio"
	"bytes"
	"github.com/herrnan/aoc-helper"
	"log"
	"strconv"
)

func main() {
	helper, err := aoc.NewHelper(aoc.Day(3), aoc.Year(2021))
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

	var nums [12]int
	for scanner.Scan() {
		text := scanner.Text()
		for pos, single := range text {
			if single == '1' {
				nums[pos]++
			} else {
				nums[pos]--
			}
		}
	}

	var res int
	for i, num := range nums {
		log.Printf("\nPos: %d\nNums: %d", i, num)

		res = res << 1
		if num > 0 {
			res = res | 0b1
		}
	}

	log.Printf("Result Binary:  %b", res)
	n := 0b111_111_111_111
	resNegate := res ^ n
	log.Printf("Result negate: %b", resNegate)

	res = res * resNegate
	log.Printf("Result: %d", res)
}

func secondPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	var in []string
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}

	oxygen := runBitFilter(in, func(num int) int {
		if num >= 0 {
			// keep 1
			return 1
		} else {
			// keep 0
			return 0
		}
	})

	log.Printf("Oxygen: %s", oxygen)

	scrubber := runBitFilter(in, func(num int) int {
		if num >= 0 {
			// keep 0
			return 0
		} else {
			// keep 1
			return 1
		}
	})

	log.Printf("Scrubber: %s", scrubber)

	oxygenInt, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	scrubberInt, err := strconv.ParseInt(scrubber, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Result: %d * %d = %d", oxygenInt, scrubberInt, oxygenInt*scrubberInt)
}

func runBitFilter(in []string, f func(int) int) string {
	for i := 0; i < 12; i++ {
		var zeroes, ones []string

		for _, s := range in {
			if s[i] == '1' {
				ones = append(ones, s)
			} else {
				zeroes = append(zeroes, s)
			}
		}

		diff := len(ones) - len(zeroes)

		if f(diff) == 1 {
			in = ones
		} else {
			in = zeroes
		}

		if len(in) == 1 {
			break
		}
	}

	log.Printf("Result: %s", in[0])
	return in[0]
}
