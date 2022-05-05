package main

import (
	"fmt"
	"math"
	"strings"
)

type Sudoku struct {
	colors          int
	n               int
	nSqrt           int
	nodesCount      int
	nodesCountSqrt  int
	adjacencyMatrix [][]int
	initialNodes    []int
	nodes           []int
}

func (s *Sudoku) Init(N, initialNode int) {
	s.n = N
	s.colors = N
	s.nSqrt = int(math.Sqrt(float64(s.n)))

	s.nodesCount = N * N
	s.nodesCountSqrt = int(math.Sqrt(float64(s.nodesCount)))

	s.initialNodes = s.getInitialNodes(initialNode)
	s.nodes = s.getInitialNodes(initialNode)

	s.initAdjacencyMatrix()
}

func (s *Sudoku) getInitialNodes(initialNode int) []int {
	initial := make([]int, s.nodesCount)
	for i := 0; i < s.nodesCount; i++ {
		initial[i] = -1
	}
	initial[initialNode] = 0
	return initial
}

func (s *Sudoku) initAdjacencyMatrix() {
	s.adjacencyMatrix = make([][]int, s.nodesCount)
	for node := range s.adjacencyMatrix {
		s.adjacencyMatrix[node] = make([]int, s.nodesCount)
	}

	s.setRowEdges()
	s.setColumnEdges()
	s.setGridEdges()
}

func (s *Sudoku) setRowEdges() {
	for i := 0; i < s.nodesCount; i++ {
		offset := int(i/s.n) * s.n

		for j := 0; j < s.n; j++ {
			column := j + offset
			if i == column {
				continue
			}
			s.adjacencyMatrix[i][column] = 1
		}
	}
}

func (s *Sudoku) setColumnEdges() {
	for i := 0; i < s.nodesCount; i++ {
		for j := i % s.n; j < s.nodesCount; j += s.n {
			if i == j {
				continue
			}
			s.adjacencyMatrix[i][j%s.nodesCount] = 1
		}
	}
}

func (s *Sudoku) setGridEdges() {
	for i := 0; i < s.nodesCount; i++ {
		// apply an offset in each odd nSqrt-sized column group
		verticalSection := int(i / (s.nSqrt * s.n))
		lineOffset := verticalSection * s.nSqrt * s.n

		horizontalSection := int(i/s.nSqrt) % s.nSqrt
		columnOffset := horizontalSection * s.nSqrt
		for j := 0; j < s.nSqrt*s.n; j += s.n {
			for k := 0; k < s.nSqrt; k++ {
				column := k + j + lineOffset + columnOffset
				if i == column {
					continue
				}
				s.adjacencyMatrix[i][column] = 1
			}
		}
	}
}

func (s *Sudoku) PrintAdjacencyMatrix() {
	fmt.Print("\t ")
	for node := range s.adjacencyMatrix {
		fmt.Printf("%d ", node%(s.nSqrt*s.n))
	}

	fmt.Println()
	for node := range s.adjacencyMatrix {
		fmt.Printf("%d\t%v\n", node, s.adjacencyMatrix[node])
	}
}

func (s *Sudoku) Solve() {
	s.solve(0)
}

func (s *Sudoku) solve(node int) bool {
	if s.isLast(node) {
		return true
	}

	for color := 0; color < s.colors; color++ {
		if s.isSafeColor(node, color) {
			s.setColor(node, color)
			if s.solve(node + 1) {
				return true
			}
		}

		if !s.isInitial(node) {
			s.nodes[node] = -1
		}
	}

	return false
}

func (s *Sudoku) isLast(node int) bool {
	return node == s.nodesCount
}

func (s *Sudoku) isSafeColor(node, color int) bool {
	if s.isInitial(node) && s.GetColor(node) == color {
		return true
	} else if s.isInitial(node) {
		return false
	}

	for i := 0; i < s.nodesCount; i++ {
		if s.AreAdjacent(node, i) && s.GetColor(i) == color {
			return false
		}
	}

	return true
}

func (s *Sudoku) isInitial(node int) bool {
	return s.initialNodes[node] != -1
}

func (s *Sudoku) GetColor(node int) int {
	return s.nodes[node]
}

func (s *Sudoku) AreAdjacent(i, j int) bool {
	return s.adjacencyMatrix[i][j] == 1
}

func (s *Sudoku) setColor(node, color int) {
	s.nodes[node] = color
}

func (s *Sudoku) Print() {
	s.printHorizontalRuler()

	for i := 0; i < s.n; i++ {
		if i%s.nSqrt == 0 {
			fmt.Println()
		}

		for j := 0; j < s.n; j++ {
			if j == 0 {
				fmt.Printf("%s  ", s.nodeToString(i+1))
			} else if j%s.nSqrt == 0 {

				fmt.Printf("  ")
			}

			fmt.Printf("%s ", s.nodeToString(s.nodes[j+i*s.n]+1))
		}
		fmt.Println()
	}
}

func (s *Sudoku) printHorizontalRuler() {
	if s.n < 10 {
		fmt.Print("//")
	} else {
		fmt.Print("///")

	}
	for j := 0; j < s.n; j++ {
		if j != 0 && j%s.nSqrt == 0 {
			fmt.Print("  ")
		}

		fmt.Printf(" %s", s.nodeToString(j+1))

	}
	fmt.Print("\n/")
}

func (s *Sudoku) nodeToString(node int) string {
	if s.n > 10 && node < 10 {
		return fmt.Sprintf(" %d", node)
	}
	return fmt.Sprint(node)
}

// Debug function
func (s *Sudoku) HasInvalidColors() bool {
	for i := 0; i < s.nodesCount; i++ {
		for j := 0; j < s.nodesCount; j++ {
			if s.AreAdjacent(i, j) && s.GetColor(i) == s.GetColor(j) {
				fmt.Println("[ERROR] Some nodes has invalid colors")
				return true
			}
		}
	}

	return false
}

// Debug function
func (s *Sudoku) HasColorlessNode() bool {
	for color := range s.nodes {
		if s.nodes[color] == -1 {
			fmt.Println("[ERROR] Has colorless nodes")
			return true
		}
	}

	return false
}

func clear() {
	fmt.Printf("\x1bc")
}

func GetBoardDimensions() int {
	var N int
	allowedValues := [3]int{4, 9, 16}

	fmt.Println("Qual o valor de N (dimensões do tabuleiro de NxN)?. Os valores permitidos são 4, 9 e 16")
	fmt.Print(">> ")
	fmt.Scan(&N)

	for v := 0; v < 3; v++ {
		if allowedValues[v] == int(N) {
			return int(N)
		}
	}

	fmt.Println("[!] Valor não permitido para N. Digite novamente")
	fmt.Println()

	return GetBoardDimensions()
}

func GetInitialNode(limit int) int {
	var node int

	fmt.Printf("Qual será a posição inicial da primeira cor/número no tabuleiro? (para o N informado, os valores permitidos vão de 1 até %d)\n", limit)
	fmt.Print(">> ")
	fmt.Scan(&node)

	node = int(node)

	if node >= 1 && node <= limit {
		return node
	}

	fmt.Println("[!] Valor não permitido. Digite novamente")
	fmt.Println()

	return GetInitialNode(limit)
}

func GetCmd() string {
	var cmd string
	fmt.Println("Digite \"q\" para sair ou \"r\" para jogar novamente.")
	fmt.Print(">> ")
	fmt.Scanf("%s", &cmd)

	if strings.ToLower(cmd) == "q" {
		return "quit"
	}

	return "run"
}

func main() {

	for cmd := "run"; cmd != "quit"; {
		clear()

		N := GetBoardDimensions()
		fmt.Print("\n\n")
		initialNode := GetInitialNode(N*N) - 1

		clear()

		var s Sudoku
		s.Init(N, initialNode)
		s.Solve()
		s.Print()

		s.HasInvalidColors()
		s.HasColorlessNode()

		fmt.Print("\n")
		fmt.Printf("Tabuleiro %dx%d", N, N)

		i := int(initialNode / s.n)
		j := int(initialNode % s.n)
		fmt.Printf("\nPosição inicial: %d (%di %dj)\n", initialNode+1, i+1, j+1)

		fmt.Print("\n\n")

		cmd = GetCmd()
	}

}
