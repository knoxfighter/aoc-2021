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
	helper, err := aoc.NewHelper(aoc.Day(4), aoc.Year(2021))
	if err != nil {
		log.Fatalln(err)
	}

	input, err := helper.GetInput()
	if err != nil {
		log.Fatalln(err)
	}

	firstPuzzle(input)
}

type Value struct {
	value int
	hit   bool
}

type Row struct {
	cols [5]Value
}

func (row *Row) addValues(_row string) {
	split := strings.Split(_row, " ")
	var i int
	for _, s := range split {
		if len(s) == 0 {
			continue
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
		row.cols[i].value = num
		i++
	}
}

func (row *Row) run(num int) {
	for i := range row.cols {
		col := &row.cols[i]
		if col.value == num {
			col.hit = true
		}
	}
}

type Field struct {
	rows [5]Row
}

func (field *Field) addRow(row string, num int) {
	field.rows[num].addValues(row)
}

func (field *Field) run(num int) {
	for i := range field.rows {
		row := &field.rows[i]
		row.run(num)
	}
}

func (field *Field) check() bool {
	var vals [5]int
	for _, row := range field.rows {
		var rVal int
		for i, col := range row.cols {
			if col.hit {
				rVal++
				vals[i]++
			}
		}

		if rVal == 5 {
			return true
		}
	}
	for _, val := range vals {
		if val == 5 {
			return true
		}
	}

	return false
}

func (field *Field) sumUnmarked() (res int) {
	for _, row := range field.rows {
		for _, col := range row.cols {
			if !col.hit {
				res += col.value
			}
		}
	}
	return
}

type Game struct {
	callNums []int
	fields   []Field
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (game *Game) run() (first int, winnerNum int, winnerField Field, last int, loserNum int, loserField Field) {
	var wons []int
	for _, num := range game.callNums {
		for i := range game.fields {
			field := &game.fields[i]
			field.run(num)
			if field.check() {
				if len(wons) == 0 {
					winnerNum = num
					first = i
					winnerField = *field
					wons = append(wons, i)
				}
				if !contains(wons, i) {
					if len(wons) == len(game.fields)-1 {
						last = i
						loserNum = num
						loserField = *field
						return
					}

					wons = append(wons, i)
				}
			}
		}
	}

	log.Fatalln("there was no winner :(")
	return
}

func firstPuzzle(input string) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	var game Game
	var block int
	var row int

	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			block++
			row = 0
			continue
		}

		if block == 0 {
			split := strings.Split(text, ",")
			for _, s := range split {
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatalln(err)
				}
				game.callNums = append(game.callNums, num)
			}
		} else {
			// read all blocks
			if row == 0 {
				//game.fields.
				game.fields = append(game.fields, Field{})
			}
			game.fields[block-1].addRow(text, row)
			row++
		}
	}

	winner, winnerNum, winnerField, loser, loserNum, loserField := game.run()

	log.Printf("winner [%d] was first to win on num [%d]", winner, winnerNum)
	winnerUnmarked := winnerField.sumUnmarked()
	log.Printf("won final score: %d", winnerNum*winnerUnmarked)

	log.Printf("loser [%d] was last to win on num [%d]", loser, loserNum)
	loserUnmarked := loserField.sumUnmarked()
	log.Printf("loser final score: %d", loserUnmarked*loserNum)
}
