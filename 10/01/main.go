package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	lines, err := input.Load("10/input.csv")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	var score int
	for _, l := range lines {
		score += scoreLine(l)
	}
	fmt.Println(score)
}

func scoreLine(line string) (score int) {
	var symbols = NewLIFO(len(line))
	var pairs = map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
	var points = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	for _, s := range strings.Split(line, "") {
		switch s {
		case "(", "[", "{", "<":
			symbols.Add(s)
		default:
			if symbols.Pop() != pairs[s] {
				return points[s] // line is corrupted.
			}
		}
	}
	return 0
}

type LIFO struct {
	values  []string
	pointer int
}

func NewLIFO(len int) LIFO {
	return LIFO{values: make([]string, len, len)}
}

func (l *LIFO) Add(s string) {
	l.values[l.pointer] = s
	l.pointer++
}

func (l *LIFO) Pop() string {
	l.pointer--
	return l.values[l.pointer]
}
