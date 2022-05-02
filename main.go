package main

import (
	"fmt"
	"math"
)

type Sudoku struct {
	n     int
	nSqrt int

	graph     [][]int
	nodes     int
	nodesSqrt int
}

func (s *Sudoku) Init(N int) {
	s.n = N
	s.nSqrt = int(math.Sqrt(float64(s.n)))
	s.nodes = N * N
	s.nodesSqrt = int(math.Sqrt(float64(s.nodes)))

	s.graph = make([][]int, s.nodes)
	for node := range s.graph {
		s.graph[node] = make([]int, s.nodes)
	}

	s.SetRowEdges()
	s.SetColumnEdges()
	s.SetGridEdges()
}

func (s *Sudoku) SetRowEdges() {
	for i := 0; i < s.nodes; i++ {
		offset := int(i/s.n) * s.n

		for j := 0; j < s.n; j++ {
			column := j + offset
			if i == column {
				continue
			}
			s.graph[i][column] = 1
		}
	}
}

func (s *Sudoku) SetColumnEdges() {
	for i := 0; i < s.nodes; i++ {
		for j := i % s.n; j < s.nodes; j += s.n {
			if i == j {
				continue
			}
			s.graph[i][j%s.nodes] = 1
		}
	}
}

func (s *Sudoku) SetGridEdges() {
	for i := 0; i < s.nodes; i++ {
		// apply an offset in each odd nSqrt-sized column group
		isOddNSizeColumn := int(i/s.nSqrt)%2 == 1
		var lineOffset int
		if isOddNSizeColumn {
			lineOffset = s.nSqrt
		}

		sectionOffset := int(i/(s.nSqrt*s.n)) * s.nSqrt * s.n

		for j := 0; j <= s.n; j += s.n {
			for k := 0; k < s.nSqrt; k++ {
				column := k + j + lineOffset + sectionOffset
				if i == column {
					continue
				}
				s.graph[i][column] = 1
			}
		}
	}
}

func (s *Sudoku) PrintMatrix() {
	fmt.Print("\t ")
	for node := range s.graph {
		fmt.Printf("%d ", node%(s.nSqrt*s.n))
	}

	fmt.Println()
	for node := range s.graph {
		fmt.Printf("%d\t%v\n", node, s.graph[node])
	}
}

func main() {
	N := 4

	var s Sudoku
	s.Init(N)
	s.PrintMatrix()
}
