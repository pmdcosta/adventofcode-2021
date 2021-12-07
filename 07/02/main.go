package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

// costsTable so we don't have to recalculate costs for the same amount of steps.
var costsTable = make(map[int]int)

func main() {
	numbers, max := getNumbers("07/input.csv")

	var lowest int

	// iterate over all possible positions.
	for i := 0; i < max; i++ {
		var fuel int
		for _, n := range numbers {
			fuel += movement(n, i)
			if fuel > lowest && lowest > 0 {
				break
			}
		}
		if fuel < lowest || lowest == 0 {
			lowest = fuel
		}
	}

	fmt.Println(lowest)
}

func movement(a, b int) int {
	distance := int(math.Abs(float64(a) - float64(b)))
	if v, ok := costsTable[distance]; ok {
		return v
	}

	var cost = 1
	var total int
	for i := 0; i < distance; i++ {
		total += cost
		cost++
	}
	costsTable[distance] = total
	return total
}

func getNumbers(s string) ([]int, int) {
	lines, err := input.Load(s)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	var max int
	var numbers []int
	for _, s := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("failed to read number", err)
		}
		numbers = append(numbers, n)
		if n > max {
			max = n
		}
	}
	return numbers, max
}
