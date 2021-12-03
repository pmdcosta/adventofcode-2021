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

	oxygen := getOxygen(lines, 0)
	o, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		log.Fatalf("failed to parse oxygen: %s", err)
	}

	co2 := getCO2(lines, 0)
	c, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		log.Fatalf("failed to parse CO2: %s", err)
	}
	fmt.Println(o, c, o*c)
}

func getOxygen(lines []string, column int) string {
	if len(lines) == 1 {
		return lines[0]
	}

	var newLines []string
	most, _ := getCommon(lines, column)
	for _, l := range lines {
		if string(l[column]) == most {
			newLines = append(newLines, l)
		}
	}

	column++
	return getOxygen(newLines, column)
}

func getCO2(lines []string, column int) string {
	if len(lines) == 1 {
		return lines[0]
	}

	var newLines []string
	_, least := getCommon(lines, column)
	for _, l := range lines {
		if string(l[column]) == least {
			newLines = append(newLines, l)
		}
	}

	column++
	return getCO2(newLines, column)
}

func getCommon(lines []string, column int) (most string, least string) {
	// check how many ones for column.
	var ones int
	for _, l := range lines {
		if l[column] == '1' {
			ones++
		}
	}

	// find the most common bit.
	if ones*2 >= len(lines) {
		return "1", "0"
	}
	return "0", "1"
}
