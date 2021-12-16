package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	board := iterateBoard(getBoard("15/input.csv"), 5)
	graph := getGraph(board)
	fmt.Println(dijkstra(graph, "0.0", fmt.Sprintf("%d.%d", len(board[0])-1, len(board)-1)))
}

// getBoard retrieves the board from the input file.
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

// iterateBoard builds the full board based on the initial tile.
func iterateBoard(board [][]int, ite int) [][]int {
	var bx, by = len(board[0]), len(board)
	var lx, ly = len(board[0]) * ite, len(board) * ite

	var bigBoard = make([][]int, ly, ly)
	for y := 0; y < ly; y++ {
		bigBoard[y] = make([]int, lx, lx)
		for x := 0; x < lx; x++ {
			var v int
			if x < bx && y < by {
				v = board[y][x]
			} else if x < bx && y >= by {
				v = bigBoard[y-by][x] + 1
				if v > 9 {
					v -= 9
				}
			} else {
				v = bigBoard[y][x-bx] + 1
				if v > 9 {
					v -= 9
				}
			}
			bigBoard[y][x] = v
		}

	}
	return bigBoard
}

// getGraph builds the adjacency map for the board.
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

// Queue list of nodes to be checked.
type Queue struct {
	values map[string]int
}

// Add adds a node to the list.
func (q *Queue) Add(node string, distance int) {
	q.values[node] = distance
}

// Pop removes the lowest distance node from the list.
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

// dijkstra finds the safest path based on the dijkstra algorithm.
func dijkstra(graph map[string]map[string]int, start, end string) int {
	visited := make(map[string]bool)             // list of visited nodes.
	dist := make(map[string]int)                 // minimum distance of each visited node to start.
	queue := Queue{values: make(map[string]int)} // list of nodes to be checked.

	// add start to the queue and start checking adjacent nodes.
	queue.Add(start, 0)
	for len(queue.values) > 0 {
		// get current node to visit.
		node := queue.Pop()

		// mark it as visited, skipping it if it's been visited before.
		if visited[node] {
			continue
		}
		visited[node] = true

		// if we reached the end, return the distance found.
		if node == end {
			return dist[end]
		}

		// check all adjacent nodes to the current one.
		for n, c := range graph[node] {
			if !visited[n] {
				// if the current path to the adjacent node is shorter than the stored distance, override it and
				// add the adjacent node to the queue.
				if dist[node]+c < dist[n] || dist[n] == 0 {
					dist[n] = dist[node] + c
					queue.Add(n, dist[node]+c)
				}
			}
		}
	}
	return dist[end]
}
