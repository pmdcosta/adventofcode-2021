package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	target := getTarget("17/input.csv")

	var hits int
	for i := -1000; i <= 1000; i++ {
		for j := -1000; j <= 1000; j++ {
			if testVelocity(Position{i, j}, target) {
				hits++
			}
		}
	}
	fmt.Println(hits)
}

func testVelocity(vel Position, target Target) bool {
	var pos = Position{0, 0}
	for i := 0; i < 1000; i++ {
		if (pos.x < target.Xmin && vel.x < 0) || (pos.x > target.Xmax && vel.x > 0) {
			return false
		}
		if pos.y < target.Ymin && vel.y <= 0 {
			return false
		}
		pos, vel = step(pos, vel)
		if onTarget(pos, target) {
			return true
		}
	}
	return false
}

func step(pos Position, velocity Position) (Position, Position) {
	pos.x = pos.x + velocity.x
	pos.y = pos.y + velocity.y

	if velocity.x > 0 {
		velocity.x--
	} else if velocity.x < 0 {
		velocity.x++
	}
	velocity.y--

	return pos, velocity
}

func onTarget(pos Position, target Target) bool {
	return (pos.x >= target.Xmin && pos.x <= target.Xmax) && (pos.y >= target.Ymin && pos.y <= target.Ymax)
}

func getTarget(file string) Target {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}
	pos := strings.Split(strings.TrimPrefix(lines[0], "target area: "), ",")
	x := strings.Split(strings.TrimPrefix(pos[0], "x="), "..")
	y := strings.Split(strings.TrimPrefix(pos[1], " y="), "..")
	x1, _ := strconv.Atoi(x[0])
	x2, _ := strconv.Atoi(x[1])
	y1, _ := strconv.Atoi(y[0])
	y2, _ := strconv.Atoi(y[1])
	return Target{x1, x2, y1, y2}
}

type Position struct {
	x int
	y int
}

type Target struct {
	Xmin int
	Xmax int
	Ymin int
	Ymax int
}
