package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	t, p := getInput("14/input.csv")

	// cycle polymerization.
	for i := 0; i < 40; i++ {
		t = polymerization(t, p)
	}

	// count number of occurrences of each letter.
	// we are considering only the first letter of each pattern, to avoid counting in duplicate.
	var out = make(map[string]int)
	for p, i := range t {
		out[strings.Split(p, "")[0]] += i
	}

	// find the letters with the highest and lowest number of occurrences.
	var max, min int
	for _, i := range out {
		if i < min || min == 0 {
			min = i
		}
		if i > max {
			max = i
		}
	}

	// return final value.
	fmt.Println(max - min)
}

func polymerization(state map[string]int, patterns map[string][]string) map[string]int {
	gen := make(map[string]int)

	// check state for known patterns.
	for s, i := range state {
		if pattern, ok := patterns[s]; ok {
			// add new patterns to the new generation.
			for _, p := range pattern {
				gen[p] += i
			}
		}
	}
	return gen
}

func getInput(file string) (template map[string]int, patterns map[string][]string) {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	// split the input string into patterns.
	template = make(map[string]int)
	temp := strings.Split(lines[0], "")
	for i := 0; i < len(temp)-1; i++ {
		template[temp[i]+temp[i+1]]++
	}

	// for each rule, map current pattern to new patterns.
	patterns = make(map[string][]string)
	for _, l := range lines[1:] {
		if l == "" {
			continue
		}
		p := strings.Split(l, " -> ")
		b := strings.Split(p[0], "")
		patterns[p[0]] = append(patterns[p[0]], b[0]+p[1], p[1]+b[1])
	}
	return template, patterns
}
