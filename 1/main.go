package main

import (
	"github.com/herrnan/aoc-helper"
	"log"
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

	log.Println(input)
}
