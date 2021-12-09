package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

const file = "09/input.csv"

func main() {
	board := readBoard(file)

	cache := make(map[Position]bool)
	lows := make(map[Position]int)
	for i := range board {
		for j := range board[i] {
			findLowPoint(board, Position{x: i, y: j}, cache, lows)
		}
	}

	var output int
	for _, l := range lows {
		output = output + l + 1
	}
	fmt.Println(output)
}

func findLowPoint(board [][]int, pos Position, cache map[Position]bool, lows map[Position]int) {
	// skip this point if we have already evaluated it.
	if _, ok := cache[pos]; ok {
		return
	}
	cache[pos] = true

	n := board[pos.x][pos.y]
	xMax := len(board) - 1
	yMax := len(board[0]) - 1

	// check if there is any lower adjacent point.
	var adjacent bool
	if pos.x != 0 && board[pos.x-1][pos.y] <= n {
		findLowPoint(board, Position{x: pos.x - 1, y: pos.y}, cache, lows)
		adjacent = true
	}
	if pos.x != xMax && board[pos.x+1][pos.y] <= n {
		findLowPoint(board, Position{x: pos.x + 1, y: pos.y}, cache, lows)
		adjacent = true
	}
	if pos.y != 0 && board[pos.x][pos.y-1] <= n {
		findLowPoint(board, Position{x: pos.x, y: pos.y - 1}, cache, lows)
		adjacent = true
	}
	if pos.y != yMax && board[pos.x][pos.y+1] <= n {
		findLowPoint(board, Position{x: pos.x, y: pos.y + 1}, cache, lows)
		adjacent = true
	}
	if adjacent {
		return
	}

	// low point found.
	lows[pos] = n
	return
}

type Position struct {
	x int
	y int
}

func readBoard(file string) (board [][]int) {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	for _, l := range lines {
		var b []int
		for _, s := range strings.Split(l, "") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("failed to read number", err)
			}
			b = append(b, n)
		}
		board = append(board, b)
	}
	return
}
