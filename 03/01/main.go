package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	lines, err := input.Load("03/input.csv")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	// check most common bit per column.
	var ones = make([]int, len(lines[0]), len(lines[0]))
	for _, l := range lines {
		for i, c := range l {
			if c == '1' {
				ones[i]++
			}
		}
	}

	// build gamma and epsilon.
	var gamma, epsilon []byte
	for _, v := range ones {
		if v*2 > len(lines) {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		} else {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		}
	}

	// convert binary number to decimal.
	g, err := strconv.ParseInt(string(gamma), 2, 64)
	if err != nil {
		log.Fatalf("failed to parse gamma: %s", err)
	}
	e, err := strconv.ParseInt(string(epsilon), 2, 64)
	if err != nil {
		log.Fatalf("failed to parse epsilon: %s", err)
	}
	fmt.Println(g * e)
}
