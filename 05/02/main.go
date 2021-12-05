package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

var rgx, _ = regexp.Compile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")

func main() {
	lines, err := input.Load("05/input.csv")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	var points int
	for _, i := range getPoints(lines) {
		if i >= 2 {
			points++
		}
	}
	fmt.Println(points)
}

func getPoints(lines []string) map[string]int {
	var world = make(map[string]int)
	for _, l := range lines {
		startX, startY, endX, endY := getStartAndEnd(l)
		if startX == endX {
			for _, y := range getValuesBetweenNumbers(startY, endY) {
				world[fmt.Sprintf("%d,%d", startX, y)]++
			}
		} else if startY == endY {
			for _, x := range getValuesBetweenNumbers(startX, endX) {
				world[fmt.Sprintf("%d,%d", x, startY)]++
			}
		} else {
			for _, s := range getValuesBetweenNumbersDiag(startX, startY, endX, endY) {
				world[s]++
			}
		}
	}
	return world
}

func getValuesBetweenNumbersDiag(ax, ay, bx, by int) []string {
	var xs []string

	var sx, sy, ey int
	if ax < bx {
		sx = ax
		sy = ay
		ey = by
	} else {
		sx = bx
		sy = by
		ey = ay
	}

	if ey > sy {
		for i := 0; i <= (ey - sy); i++ {
			xs = append(xs, fmt.Sprintf("%d,%d", sx+i, sy+i))
		}
	} else {
		for i := 0; i <= (sy - ey); i++ {
			xs = append(xs, fmt.Sprintf("%d,%d", sx+i, sy-i))
		}
	}
	return xs
}

func getValuesBetweenNumbers(a, b int) []int {
	var xs []int
	for i := a; i <= b; i++ {
		xs = append(xs, i)
	}
	for i := b; i <= a; i++ {
		xs = append(xs, i)
	}
	return xs
}

func getStartAndEnd(l string) (int, int, int, int) {
	var p = rgx.FindStringSubmatch(l)
	return aToi(p[1]), aToi(p[2]), aToi(p[3]), aToi(p[4])
}

func aToi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("failed to read number", err)
	}
	return n
}
