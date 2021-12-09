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

	// check if there is any lower adjacent point.
	var adjacent bool
	for _, p := range []Position{{pos.x - 1, pos.y}, {pos.x + 1, pos.y}, {pos.x, pos.y - 1}, {pos.x, pos.y + 1}} {
		if p.x >= 0 && p.x <= len(board)-1 && p.y >= 0 && p.y <= len(board[0])-1 {
			if board[p.x][p.y] <= board[pos.x][pos.y] {
				findLowPoint(board, p, cache, lows)
				adjacent = true
			}
		}
	}
	if adjacent {
		return
	}

	// low point found.
	lows[pos] = board[pos.x][pos.y]
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
