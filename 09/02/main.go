package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

const file = "09/input.csv"

func main() {
	board := readBoard(file)

	var cache = make(map[Position]Position)
	var lows = make(map[Position]Set)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 9 {
				continue // ignore heights of 9.
			}
			findLowPoint(board, Position{x: i, y: j}, Set{}, cache, lows)
		}
	}

	// get the 3 biggest basins.
	var basins []int
	for _, b := range lows {
		basins = append(basins, len(b))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Println(basins[0] * basins[1] * basins[2])
}

func findLowPoint(board [][]int, pos Position, steps Set, cache map[Position]Position, lows map[Position]Set) {
	// check if this position is already in the current path, to avoid loops.
	if _, ok := steps[pos]; ok {
		return
	}
	steps[pos] = true

	// skip this point if we have already evaluated it.
	// if it points to a low, then add the current path to the low basin.
	if c, ok := cache[pos]; ok {
		if _, ok := lows[c]; ok {
			for s := range steps {
				lows[c][s] = true // add all steps to the low.
				cache[s] = c      // update all steps in path to point to low.
			}
		}
		return
	}

	// check if there is any lower adjacent point.
	var adjacent bool
	for _, p := range []Position{{pos.x - 1, pos.y}, {pos.x + 1, pos.y}, {pos.x, pos.y - 1}, {pos.x, pos.y + 1}} {
		if p.x >= 0 && p.x <= len(board)-1 && p.y >= 0 && p.y <= len(board[0])-1 {
			if board[p.x][p.y] <= board[pos.x][pos.y] {
				findLowPoint(board, p, steps, cache, lows)
				adjacent = true
			}
		}
	}
	if adjacent {
		return
	}

	// low point found.
	lows[pos] = make(Set)
	for s := range steps {
		lows[pos][s] = true
		cache[s] = pos
	}
	return
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

type Position struct {
	x int
	y int
}

type Set map[Position]bool
