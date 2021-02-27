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
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func MatrixMultiplication(matrix1, matrix2 [][]float64) ([][]float64, error) {
	if len(matrix1) == 0 || len(matrix2) == 0 {
		return nil, errors.New("Nil Matrix Length")
	}

	matr