package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	lines, err := input.Load("07/input.csv")
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

	var minDistance int

	// iterate over all possible positions.
	for i := 0; i < max; i++ {
		var cumulative int
		for _, n := range numbers {
			cumulative += int(math.Abs(float64(n) - float64(i)))
			if cumulative > minDistance && minDistance > 0 {
				break
			}
		}
		if cumulative < minDistance || minDistance == 0 {
			minDistance = cumulative
		}
	}
	fmt.Println(minDistance)
}
