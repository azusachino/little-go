package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type Point struct {
	i, j int
}

func (p Point) add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

var directions = [4]Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end Point) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []Point{start}

	for len(Q) > 0 {
		// cur := Q[0]
	}
}

func main() {
	maze := readMaze("maze/maze.in")
	fmt.Println(maze)

	walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1})
}
