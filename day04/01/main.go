package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed input.txt
var data string

type direction struct {
	x int
	y int
}

// DIRECTIONS y,x -axis. Coordinats are in the right lower quadrant of a coordinat diagram (south is positive)!
var NE = direction{y: -1, x: 1}
var N = direction{y: -1, x: 0}
var NW = direction{y: -1, x: -1}
var W = direction{y: 0, x: -1}
var SW = direction{y: 1, x: -1}
var S = direction{y: 1, x: 0}
var SE = direction{y: 1, x: 1}
var E = direction{y: 0, x: 1}

func main() {
	// 	data := `MMMSXXMASM
	// MSAMXMSMSA
	// AMXSXMAAMM
	// MSAMASMSMX
	// XMASAMXAMM
	// XXAMMXXAMA
	// SMSMSASXSS
	// SAXAMASAAA
	// MAMMMXMMMM
	// MXMXAXMASX`

	// Build the matrix
	matrix, err := buildMatrix(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find 'X' as starting points
	// From every starting Point: Traverse in all directions and count "XMAS"
	xmas := findXmas(matrix)

	fmt.Println("The Sum of 'XMAS' is:", xmas)
}

func buildMatrix(data string) ([][]byte, error) {
	// 140x140 Matrix
	matrix := make([][]byte, 140) // 1st dimension: Y-Axis; 2nd dimension: X-Axis
	for y := range matrix {
		matrix[y] = make([]byte, 140)
	}
	scanner := bufio.NewScanner(strings.NewReader(data))

	// line by line
	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()
		if len(text) > 140 {
			return nil, errors.New("scanned line is to large! Max length 140")
		}

		matrix[y] = []byte(text)
	}
	return matrix, nil
}

func findXmas(matrix [][]byte) int {
	xmasSum := 0

	for y, row := range matrix {
		for x, char := range row {
			if rune(char) == rune('X') {
				xmasSum += findXmasAtCoordinate(matrix, y, x)
			}
		}
	}

	return xmasSum
}

func findXmasAtCoordinate(matrix [][]byte, y int, x int) int {
	xmasSum := 0
	xmasSum += traverseMatrix(matrix, y, x, NE)
	xmasSum += traverseMatrix(matrix, y, x, N)
	xmasSum += traverseMatrix(matrix, y, x, NW)
	xmasSum += traverseMatrix(matrix, y, x, W)
	xmasSum += traverseMatrix(matrix, y, x, SW)
	xmasSum += traverseMatrix(matrix, y, x, S)
	xmasSum += traverseMatrix(matrix, y, x, SE)
	xmasSum += traverseMatrix(matrix, y, x, E)
	return xmasSum
}

func traverseMatrix(matrix [][]byte, y int, x int, dir direction) int {
	xmasBuffer := []byte{'M', 'A', 'S'} // Location of 'X' already known

	// check chars
	for i := 0; i < 3; i++ {
		yTarget := y + dir.y
		xTarget := x + dir.x

		// out of bounds safety
		if yTarget < 0 || yTarget >= len(matrix) || xTarget < 0 || xTarget >= len(matrix[0]) {
			return 0
		}

		if matrix[yTarget][xTarget] != xmasBuffer[i] {
			return 0 // return on mismatch
		}
		// increment or decrement based on sign. Stop if 0 in a direction
		if dir.y != 0 {
			dir.y += dir.y / abs(dir.y)
		}
		if dir.x != 0 {
			dir.x += dir.x / abs(dir.x)
		}

	}
	return 1 // found xmas! Add up 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
