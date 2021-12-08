package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

const file = "08/input.csv"

func main() {
	entries := getEntries(file)

	var values int

	for _, e := range entries {
		for _, o := range e.output {
			switch len(o) {
			case 2, 3, 4, 7:
				values++
			}
		}
	}
	fmt.Println(values)

}

func getEntries(file string) (entries []Entry) {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	for _, l := range lines {
		var e Entry
		line := strings.Split(l, "|")
		patterns := strings.Split(line[0], " ")
		for _, p := range patterns {
			e.patterns = append(e.patterns, p)
		}
		output := strings.Split(line[1], " ")
		for _, o := range output {
			e.output = append(e.output, o)
		}
		entries = append(entries, e)
	}
	return entries
}

type Entry struct {
	patterns []string
	output   []string
}
