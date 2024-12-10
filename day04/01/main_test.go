package main

import (
	"testing"
)

func TestBuildMatrix(t *testing.T) {
	// 3x3 matrix -> walk from the '5' in every direction
	data := `123
456
789`

	matrix, err := buildMatrix(data)
	if err != nil {
		t.Error(err)
	}

	if matrix == nil {
		t.Error("Matrix was not build")
	}
	if len(matrix[0]) != 3 {
		t.Error("Rows of incorrect size! Size was:", len(matrix[0]))
	}
}
