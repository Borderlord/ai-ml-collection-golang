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
	out := make([][]float32, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(