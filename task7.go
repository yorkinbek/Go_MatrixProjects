package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Printf("n=")
	fmt.Scanf("%d", &n)

	fmt.Printf("m=")
	fmt.Scanf("%d", &m)

	matrix1 := make([][]float64, n)
	createMatrix(matrix1, n, m)

	matrix2 := make([][]float64, n)
	createMatrix(matrix2, n, m)

	addMatrix(matrix1, matrix2, n, m)
	fmt.Println("ADD: matrix1 * matrix2: ", addMatrix(matrix1, matrix2, n, m))

	MulMatrix(matrix1, matrix2, n, m)
	fmt.Println("MUL: matrix1 * matrix2: ", MulMatrix(matrix1, matrix2, n, m))

}

func addMatrix(matrix1 [][]float64, matrix2 [][]float64, n, m int) [][]float64 {
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			result[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}
	return result
}

func MulMatrix(matrix1 [][]float64, matrix2 [][]float64, n, m int) [][]float64 {
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}

func createMatrix(matrix [][]float64, n, m int) [][]float64 {
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			fmt.Printf("matrix[%d][%d] : ", i, j)
			fmt.Scanf("%f", &matrix[i][j])
		}
	}
	return matrix
}
