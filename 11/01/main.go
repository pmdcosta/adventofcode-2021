package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	world := getWorld("11/input.csv")

	var flashes int
	for i := 0; i < 100; i++ {
		flashes += step(world)
	}
	fmt.Println(flashes)
}

func step(world [][]int) (flashes int) {
	var flashed = make(map[string]bool)

	// increase every octopus.
	for x, l := range world {
		for y := range l {
			increase(world, x, y, flashed)
		}
	}

	// reset energy when flashing.
	for x, l := range world {
		for y, o := range l {
			if o > 9 {
				world[x][y] = 0
				flashes++
			}
		}
	}

	return flashes
}

func increase(world [][]int, x int, y int, flashed map[string]bool) {
	if _, ok := flashed[fmt.Sprintf("%d.%d", x, y)]; ok {
		return
	}

	// stay inside the world.
	if x < 0 || x == len(world) || y < 0 || y == len(world[0]) {
		return
	}

	// increase in energy.
	world[x][y] = world[x][y] + 1

	// if a flash happens, increase every nearby octopus.
	if world[x][y] > 9 {
		flashed[fmt.Sprintf("%d.%d", x, y)] = true
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				increase(world, i, j, flashed)
			}
		}
	}
}

func getWorld(file string) (world [][]int) {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	for _, l := range lines {
		var b []int
		for _, c := range strings.Split(l, "") {
			n, _ := strconv.Atoi(c)
			b = append(b, n)
		}
		world = append(world, b)
	}
	return world
}
