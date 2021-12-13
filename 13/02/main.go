package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

type Position struct {
	x, y int
}

func main() {
	board, vertical, horizontal := getWorld("13/input.csv")

	var result = make(map[Position]bool)
	for p := range board {
		np := fold(p, vertical, horizontal)
		result[np] = true
	}

	printWorld(result)
	fmt.Println(len(result))
}

func getWorld(file string) (board map[Position]bool, vertical []int, horizontal []int) {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	board = make(map[Position]bool)
	for _, l := range lines {
		if l == "" {
			continue
		} else if strings.Contains(l, "fold along y=") {
			y, _ := strconv.Atoi(strings.TrimPrefix(l, "fold along y="))
			vertical = append(vertical, y)
		} else if strings.Contains(l, "fold along x=") {
			x, _ := strconv.Atoi(strings.TrimPrefix(l, "fold along x="))
			horizontal = append(horizontal, x)
		} else if p := strings.Split(l, ","); len(p) == 2 {
			n1, _ := strconv.Atoi(p[0])
			n2, _ := strconv.Atoi(p[1])
			board[Position{n1, n2}] = true
		}
	}
	return board, vertical, horizontal
}

func printWorld(board map[Position]bool) {
	var x, y int
	for p := range board {
		if p.x > x {
			x = p.x
		}
		if p.y > y {
			y = p.y
		}
	}
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			if _, ok := board[Position{j, i}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func fold(p Position, vertical, horizontal []int) Position {
	for _, v := range vertical {
		if p.y > v {
			p.y = p.y - 2*(p.y-v)
		}
	}
	for _, h := range horizontal {
		if p.x > h {
			p.x = p.x - 2*(p.x-h)
		}
	}
	return p
}
