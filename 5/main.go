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
	helper, err := aoc.NewHelper(aoc.Day(5), aoc.Year(2021))
	if err != nil {
		log.Fatalln(err)
	}

	input, err := helper.GetInput()
	if err != nil {
		log.Fatalln(err)
	}

	firstPuzzle(input)
}

type Pos struct {
	x int
	y int
}

// two-dimensional dynamic array
var field map[int]map[int]int

func firstPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)
	field = make(map[int]map[int]int)

	for scanner.Scan() {
		text := scanner.Text()

		positions := strings.Split(text, " -> ")
		fromPosition := strings.Split(positions[0], ",")
		toPosition := strings.Split(positions[1], ",")

		var err error
		var fromPos Pos
		fromPos.x, err = strconv.Atoi(fromPosition[0])
		if err != nil {
			log.Fatalln(err)
		}
		fromPos.y, err = strconv.Atoi(fromPosition[1])
		if err != nil {
			log.Fatalln(err)
		}

		var toPos Pos
		toPos.x, err = strconv.Atoi(toPosition[0])
		if err != nil {
			log.Fatalln(err)
		}
		toPos.y, err = strconv.Atoi(toPosition[1])
		if err != nil {
			log.Fatalln(err)
		}

		if fromPos.x == toPos.x {
			if fromPos.y < toPos.y {
				for i := fromPos.y; i <= toPos.y; i++ {
					if field[fromPos.x] == nil {
						field[fromPos.x] = make(map[int]int)
					}
					field[fromPos.x][i]++
				}
			} else if fromPos.y > toPos.y {
				for i := toPos.y; i <= fromPos.y; i++ {
					if field[fromPos.x] == nil {
						field[fromPos.x] = make(map[int]int)
					}
					field[fromPos.x][i]++
				}
			} else {
				log.Printf("thats bad :(\nfrom %v to %v", fromPos, toPos)
			}
		} else if fromPos.y == toPos.y {
			if fromPos.x < toPos.x {
				for i := fromPos.x; i <= toPos.x; i++ {
					if field[i] == nil {
						field[i] = make(map[int]int)
					}
					field[i][fromPos.y]++
				}
			} else if fromPos.x > toPos.x {
				for i := toPos.x; i <= fromPos.x; i++ {
					if field[i] == nil {
						field[i] = make(map[int]int)
					}
					field[i][fromPos.y]++
				}
			} else {
				log.Printf("thats bad2 :(\nfrom %v to %v", fromPos, toPos)
			}
		} else {
			if fromPos.x < toPos.x && fromPos.y < toPos.y {
				for pos := fromPos; pos.x <= toPos.x; pos.x, pos.y = pos.x+1, pos.y+1 {
					if field[pos.x] == nil {
						field[pos.x] = make(map[int]int)
					}
					field[pos.x][pos.y]++
				}
			} else if fromPos.x > toPos.x && fromPos.y < toPos.y {
				for pos := fromPos; pos.x >= toPos.x; pos.x, pos.y = pos.x-1, pos.y+1 {
					if field[pos.x] == nil {
						field[pos.x] = make(map[int]int)
					}
					field[pos.x][pos.y]++
				}
			} else if fromPos.x < toPos.x && fromPos.y > toPos.y {
				for pos := fromPos; pos.x <= toPos.x; pos.x, pos.y = pos.x+1, pos.y-1 {
					if field[pos.x] == nil {
						field[pos.x] = make(map[int]int)
					}
					field[pos.x][pos.y]++
				}
			} else if fromPos.x > toPos.x && fromPos.y > toPos.y {
				for pos := fromPos; pos.x >= toPos.x; pos.x, pos.y = pos.x-1, pos.y-1 {
					if field[pos.x] == nil {
						field[pos.x] = make(map[int]int)
					}
					field[pos.x][pos.y]++
				}
			} else {
				log.Printf("thats bad 3 :(")
			}
		}
	}

	var dangerous int
	for _, subfield := range field {
		for _, point := range subfield {
			if point >= 2 {
				dangerous++
			}
		}
	}

	log.Printf("Dangerous points: %d", dangerous)
}
