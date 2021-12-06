package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	lines, err := input.Load("06/test.csv")
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

	// iterate over all generations.
	for i := 0; i < 256; i++ {
		for i, f := range numbers {
			f--
			if f == -1 {
				f = 6
				numbers = append(numbers, 8)
			}
			numbers[i] = f
		}
	}
	fmt.Println(len(numbers))
}
