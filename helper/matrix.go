package helper

import (
	"errors"
)

func MatrixAdditionWithNumber(matrix [][]float64, number float64) [][]float64 {
	for row := range matrix {
		for col := range matrix[row] {
			matrix[row][col] += number
		}
	}
	return matrix
}

func TransponseMatrix(x [][]float32) [][]float32 {
	out := make([][]float32, 