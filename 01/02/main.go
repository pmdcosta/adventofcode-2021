package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// get current path.
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// read file.
	depths, err := readDepths(filepath.Join(path, "01/input.csv"))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read depths file: %w", err))
	}

	// calculate increases.
	var last, increases int
	for i := 2; i < len(depths); i++ {
		d := depths[i-2] + depths[i-1] + depths[i]
		if d > last {
			increases++
		}
		last = d
	}
	fmt.Println(increases - 1) // skip the first line.
}

func readDepths(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file lines: %w", err)
	}

	var out = make([]int, 0, len(lines))
	for _, l := range lines {
		intL, err := strconv.Atoi(l)
		if err != nil {
			return nil, fmt.Errorf("failed to convert line to int: %w", err)
		}
		out = append(out, intL)
	}
	return out, nil
}
