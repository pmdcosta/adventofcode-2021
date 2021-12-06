package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

const (
	states      = 9   // number of possible timer states.
	generations = 256 // generations to iterate.
)

func main() {
	lines, err := input.Load("06/input.csv")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	var numbers []int
	for _, s := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("failed to read number", err)
		}
		numbers = append(numbers, n)
	}

	// build population histogram.
	var population = make([]int, states, states)
	for i := 0; i < len(numbers); i++ {
		population[numbers[i]]++
	}

	// update population histogram over generations.
	for i := 0; i < generations; i++ {
		var gen = make([]int, states, states)
		gen[0] = population[1]                 // previous gen 1 become new gen 0.
		gen[1] = population[2]                 // previous gen 2 become new gen 1.
		gen[2] = population[3]                 // previous gen 3 become new gen 2.
		gen[3] = population[4]                 // previous gen 4 become new gen 3.
		gen[4] = population[5]                 // previous gen 5 become new gen 4.
		gen[5] = population[6]                 // previous gen 6 become new gen 5.
		gen[6] = population[0] + population[7] // previous gen 0 and gen 7 become new gen 6.
		gen[7] = population[8]                 // previous gen 8 become new gen 7.
		gen[8] = population[0]                 // previous gen 0 become create new gen 8.
		population = gen
	}

	var fish int
	for _, p := range population {
		fish += p
	}
	fmt.Println(fish)
}
