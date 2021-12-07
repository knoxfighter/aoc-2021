package main

import (
	"bufio"
	"bytes"
	"github.com/herrnan/aoc-helper"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	log.Printf("in nano: %d", elapsed.Nanoseconds())
}

func main() {
	start := time.Now()

	helper, err := aoc.NewHelper(aoc.Day(7), aoc.Year(2021))
	if err != nil {
		log.Fatalln(err)
	}

	input, err := helper.GetInput()
	if err != nil {
		log.Fatalln(err)
	}

	timeTrack(start, "downloading")

	firstPuzzle(input)
	secondPuzzle(input)
}

func firstPuzzle(input string) {
	start := time.Now()

	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	var data []int

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalln(err)
			}

			data = append(data, i)
		}
	}

	sort.Ints(data)

	var val int
	if len(data)%2 == 0 {
		val = data[len(data)/2]
	} else {
		midValue := len(data) / 2
		data1 := data[midValue]
		data2 := data[midValue+1]
		data1 += data2
		val = data1 / 2
	}

	log.Printf("Median: %d", val)

	var sprit int
	for _, datum := range data {
		if datum < val {
			sprit += val - datum
		} else {
			sprit += datum - val
		}
	}

	log.Printf("used sprit: %d", sprit)

	timeTrack(start, "firstPuzzle")
}

func secondPuzzle(input string) {
	start := time.Now()

	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	var data []int

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalln(err)
			}

			data = append(data, i)
		}
	}

	var val int
	for _, datum := range data {
		val += datum
	}

	f := float64(val) / float64(len(data))
	f = math.Floor(f)
	val = int(f)

	log.Printf("middle value: %d", val)

	var sprit int
	for _, datum := range data {
		if datum < val {
			diff := val - datum
			sprit += (diff * (diff + 1)) / 2
		} else {
			diff := datum - val
			sprit += (diff * (diff + 1)) / 2
		}
	}

	log.Printf("used sprit: %d", sprit)

	timeTrack(start, "secondPuzzle")
}
