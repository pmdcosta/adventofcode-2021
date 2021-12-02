package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Direction string

const (
	directionUp      Direction = "up"
	directionDown    Direction = "down"
	directionForward Direction = "forward"
)

type movement struct {
	direction Direction
	value     int
}

func main() {
	directions, err := readDirections()
	if err != nil {
		log.Fatal(err)
	}

	var x, y, aim int
	for _, d := range directions {
		switch d.direction {
		case directionForward:
			x += d.value
			y = y + aim*d.value
		case directionDown:
			aim += d.value
		case directionUp:
			aim -= d.value
		}
	}
	fmt.Println(x * y)
}

func readDirections() (out []movement, err error) {
	input, err := readInputFile()
	if err != nil {
		return nil, err
	}
	for _, l := range input {
		line := strings.Split(l, " ")
		val, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, fmt.Errorf("failed to convert line value to int: %w", err)
		}
		out = append(out, movement{direction: Direction(line[0]), value: val})
	}
	return out, nil
}

const inputFile = "02/input.csv"

func readInputFile() (lines []string, err error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current path: %w", err)
	}
	file, err := os.Open(filepath.Join(path, inputFile))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file lines: %w", err)
	}
	return lines, nil
}
