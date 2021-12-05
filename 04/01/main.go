package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	lines, err := input.Load("04/input.csv")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	numbers := getNumbers(lines[0])
	boards := getBoards(lines[1:])

	// range through all numbers.
	for _, n := range numbers {
		for _, b := range boards {
			if won, score := checkBoard(b, n); won {
				fmt.Println(score)
				return
			}
		}
	}
}

func checkBoard(b board, number int) (won bool, score int) {
	for i, l := range b.numbers {
		for j, n := range l {
			if n == number {
				b.lines[i]++
				b.columns[j]++
				b.numbers[i][j] = -1
				if b.lines[i] == 5 || b.columns[j] == 5 {
					return true, getScore(b.numbers, number)
				}
			}
		}
	}
	return false, 0
}

func getScore(numbers [][]int, last int) (score int) {
	for _, l := range numbers {
		for _, n := range l {
			if n != -1 {
				score += n
			}
		}
	}
	return score * last
}

func getNumbers(line string) (numbers []int) {
	for _, n := range strings.Split(line, ",") {
		v, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("failed to read number", err)
		}
		numbers = append(numbers, v)
	}
	return numbers
}

type board struct {
	numbers [][]int
	lines   []int
	columns []int
}

func getBoards(lines []string) []board {
	// read all numbers.
	var numbers []int
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		for _, n := range strings.Split(l, " ") {
			if strings.TrimSpace(n) == "" {
				continue
			}
			n, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal("failed to read number ", err)
			}
			numbers = append(numbers, n)
		}
	}

	// group numbers into boards.
	var boards []board
	for i := 0; i < len(numbers); i += 25 {
		boards = append(boards, getBoard(numbers[i:i+25]))
	}
	return boards
}

func getBoard(numbers []int) board {
	var b [][]int
	for i := 0; i < len(numbers); i += 5 {
		b = append(b, numbers[i:i+5])
	}
	return board{numbers: b, lines: make([]int, 5, 5), columns: make([]int, 5, 5)}
}
