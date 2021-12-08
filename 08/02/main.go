package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

const file = "08/input.csv"

// Entry represents an input line.
type Entry struct {
	patterns []Pattern
	output   []string
}

// Pattern is a group of segments that represents a digit.
type Pattern struct {
	set map[string]bool
}

// Add adds a string to a pattern.
func (p *Pattern) Add(s string) {
	if p.set == nil {
		p.set = make(map[string]bool)
	}
	p.set[s] = true
}

// Contains checks whether the pattern contains a string.
func (p *Pattern) Contains(s string) bool {
	_, ok := p.set[s]
	return ok
}

// String returns a sorted string representation of a pattern.
func (p *Pattern) String() string {
	var s []string
	for l := range p.set {
		s = append(s, l)
	}
	sort.Strings(s)
	return strings.Join(s, "")
}

// sortStr sorts a string.
func sortStr(o string) string {
	var ds = strings.Split(o, "")
	sort.Strings(ds)
	return strings.Join(ds, "")
}

// del removes an element from the slice.
func del(values []Pattern, i int) []Pattern {
	var newValues []Pattern
	for j, v := range values {
		if i != j {
			newValues = append(newValues, v)
		}
	}
	return newValues
}

// getInput reads the input into Entries.
func getInput(file string) (entries []Entry) {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	for _, l := range lines {
		var e = Entry{patterns: make([]Pattern, 0)}

		// read patterns.
		line := strings.Split(l, "|")
		patterns := strings.Split(line[0], " ")
		for _, pattern := range patterns {
			if pattern == "" {
				continue
			}
			var p Pattern
			for _, v := range pattern {
				p.Add(fmt.Sprintf("%c", v))
			}
			e.patterns = append(e.patterns, p)
		}

		// read output
		output := strings.Split(line[1], " ")
		for _, o := range output {
			if o == "" {
				continue
			}
			e.output = append(e.output, sortStr(o))
		}
		entries = append(entries, e)
	}
	return entries
}

func main() {
	var sum int
	for _, e := range getInput(file) {
		sum += e.getOutputCode()
	}
	fmt.Println(sum)
}

// getOutputCode calculates the output code by comparing the output with the digit mappings.
func (e Entry) getOutputCode() (code int) {
	digits := e.process()
	var codeStr []string
	for _, o := range e.output {
		for i, d := range digits {
			if o == d.String() {
				codeStr = append(codeStr, fmt.Sprintf("%v", i))
			}
		}
	}
	code, _ = strconv.Atoi(strings.Join(codeStr, ""))
	return code
}

// process maps patterns to the corresponding digit.
func (e *Entry) process() []Pattern {
	var digits = make([]Pattern, 10, 10)
	var aS, cS string

	// group digits based on length.
	var len5, len6 []Pattern
	for _, p := range e.patterns {
		switch len(p.set) {
		case 2:
			digits[1] = p
		case 3:
			digits[7] = p
		case 4:
			digits[4] = p
		case 5:
			len5 = append(len5, p)
		case 6:
			len6 = append(len6, p)
		case 7:
			digits[8] = p
		}
	}

	// step 1: find value that 7 contains and 1 doesn't.
	for s := range digits[7].set {
		if !digits[1].Contains(s) {
			aS = s
		}
	}

	// step 2: find value that 7 contains and 6 doesn't.
	for v := range digits[7].set {
		if v == aS {
			continue
		}
		for i, d := range len6 {
			if !d.Contains(v) {
				len6 = del(len6, i)
				digits[6] = d
				cS = v
			}
		}
	}

	// step 3: find value that 4 contains and 0 doesn't.
	// We can also infer 9 as it's the only leftover digit with 6 segments.
	for v := range digits[4].set {
		for i, d := range len6 {
			if !d.Contains(v) {
				digits[0] = d
				len6 = del(len6, i)
				digits[9] = len6[0]
			}
		}
	}

	// step 4: find leftover digit that contains 1.
	for i, d := range len5 {
		var match = true
		for v := range digits[1].set {
			if !d.Contains(v) {
				match = false
			}
		}
		if match {
			digits[3] = d
			len5 = del(len5, i)
		}
	}

	// step 5: find leftover digit contains C.
	for i, d := range len5 {
		if !d.Contains(cS) {
			digits[5] = d
			len5 = del(len5, i)
			digits[2] = len5[0]
		}
	}

	return digits
}
