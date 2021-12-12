package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

var paths int

func main() {
	caves := getCaves("12/input.csv")
	traverseCave(caves["start"], make([]string, 0))
	fmt.Println(paths)
}

func traverseCave(cave Cave, path []string) {
	// add cave to current path.
	path = append(path, cave.Name)

	// path has reached the end.
	if cave.Name == "end" {
		paths++
		return
	}

	// recursively traverse caves, skipping the ones we already visited.
	for n, c := range cave.Connections {
		if !visited(path, n) || !c.Small {
			traverseCave(c, path)
		}
	}
}

// visited checks if we have traversed the requested cave already.
func visited(path []string, name string) bool {
	for _, s := range path {
		if s == name {
			return true
		}
	}
	return false
}

type Cave struct {
	Name        string
	Small       bool
	Connections map[string]Cave
}

func NewCave(name string, small bool) Cave {
	return Cave{Name: name, Small: small, Connections: make(map[string]Cave)}
}

func getCaves(file string) map[string]Cave {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	var caves = make(map[string]Cave)
	for _, l := range lines {
		var c1, c2 Cave
		p := strings.Split(l, "-")
		if c, ok := caves[p[0]]; ok {
			c1 = c
		} else {
			c1 = NewCave(p[0], strings.ToUpper(p[0]) != p[0])
		}
		if c, ok := caves[p[1]]; ok {
			c2 = c
		} else {
			c2 = NewCave(p[1], strings.ToUpper(p[1]) != p[1])
		}
		c1.Connections[p[1]] = c2
		c2.Connections[p[0]] = c1
		caves[p[0]] = c1
		caves[p[1]] = c2
	}

	return caves
}
