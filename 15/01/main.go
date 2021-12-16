package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	board := getBoard("15/test.csv")
	graph := getGraph(board)
	fmt.Println(dijkstra(graph, "0.0", fmt.Sprintf("%d.%d", len(board[0])-1, len(board)-1)))
}

func getBoard(file string) (board [][]int) {
	lines, err := input.Load(file)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	for _, l := range lines {
		var line []int
		for _, n := range strings.Split(l, "") {
			v, _ := strconv.Atoi(n)
			line = append(line, v)
		}
		board = append(board, line)
	}
	return board
}

func getGraph(board [][]int) map[string]map[string]int {
	var graph = make(map[string]map[string]int)
	for y, l := range board {
		for x := range l {
			n := fmt.Sprintf("%d.%d", x, y)
			graph[n] = make(map[string]int)
			if y+1 < len(board) {
				graph[n][fmt.Sprintf("%d.%d", x, y+1)] = board[y+1][x]
			}
			if y-1 >= 0 {
				graph[n][fmt.Sprintf("%d.%d", x, y-1)] = board[y-1][x]
			}
			if x+1 < len(board) {
				graph[n][fmt.Sprintf("%d.%d", x+1, y)] = board[y][x+1]
			}
			if x-1 >= 0 {
				graph[n][fmt.Sprintf("%d.%d", x-1, y)] = board[y][x-1]
			}
		}
	}
	return graph
}

type Queue struct {
	values map[string]int
}

func (q *Queue) Add(node string, distance int) {
	q.values[node] = distance
}

func (q *Queue) Pop() string {
	var dist int
	var node string
	for n, d := range q.values {
		if dist == 0 || d < dist {
			dist = d
			node = n
		}
	}
	delete(q.values, node)
	return node
}

func dijkstra(graph map[string]map[string]int, start, end string) int {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	queue := Queue{values: make(map[string]int)}

	dist[start] = 0
	queue.Add(start, 0)

	for len(queue.values) > 0 {
		node := queue.Pop()

		if visited[node] {
			continue
		}
		visited[node] = true

		for n, c := range graph[node] {
			if !visited[n] {
				if dist[node]+c < dist[n] || dist[n] == 0 {
					dist[n] = dist[node] + c
					queue.Add(n, dist[node]+c)
				}
			}
		}
	}
	return dist[end]
}
