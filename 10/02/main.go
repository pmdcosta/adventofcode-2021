package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	lines, err := input.Load("10/input.csv")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	var scores []int
	for _, l := range lines {
		if s, err := evaluateLine(l); err == nil {
			scores = append(scores, s)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func evaluateLine(line string) (int, error) {
	var symbols = NewLIFO(len(line))
	var pairs = map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
	for _, s := range strings.Split(line, "") {
		switch s {
		case "(", "[", "{", "<":
			symbols.Add(s)
		default:
			if symbols.Pop() != pairs[s] {
				return 0, errors.New("line is corrupted")
			}
		}
	}
	return getScore(completeLine(symbols)), nil
}

func completeLine(symbols LIFO) []string {
	var pairs = map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
	var missing []string
	for _, m := range symbols.PopAll() {
		missing = append(missing, pairs[m])
	}
	return missing
}

func getScore(val []string) (score int) {
	var points = map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	for _, v := range val {
		score = score*5 + points[v]
	}
	return score
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

func (l *LIFO) PopAll() []string {
	var values []string
	for i := l.pointer - 1; i >= 0; i-- {
		values = append(values, l.values[i])
	}
	l.pointer = 0
	return values
}
