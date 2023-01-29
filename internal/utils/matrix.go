package utils

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

type Matrix [][]int

func IsQuadratic(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}

func Modulo(a, m int) int {
	result := a % m
	if (result < 0 && m > 0) || (result > 0 && m < 0) {
		return result + m
	}
	return result
}

func ModInverse(a, m int) int {
	a = Modulo(a, m)
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return -1
}

func InitializeMatrix(row, col int) Matrix {
	matrix := make(Matrix, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
	}
	return matrix
}

func GenerateKeyMatrix(key string, size int) Matrix {
	matrix := InitializeMatrix(size, size)

	k := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = int(key[k]) - 65
			k++
		}
	}
	return matrix
}

func GenerateSegmentedText(plaintext string, size int) []Matrix {
	numOfSegments := len(plaintext) / size

	segmentedPlaintText := make([]Matrix, numOfSegments)

	k := 0
	for i := 0; i < numOfSegments; i++ {
		matrix := InitializeMatrix(size, 1)
		for j := 0; j < size; j++ {
			matrix[j][0] = int(plaintext[k]) - 65
			k++
		}
		segmentedPlaintText[i] = matrix
	}
	return segmentedPlaintText
}

func (matrix *Matrix) Multiply(matrixB Matrix) Matrix {
	rowA, colA := len(*matrix), len((*matrix)[0])
	colB := len(matrixB[0])

	result := InitializeMatrix(rowA, colB)

	for i := 0; i < rowA; i++ {
		for j := 0; j < colB; j++ {
			for k := 0; k < colA; k++ {
				result[i][j] += (*matrix)[i][k] * matrixB[k][j]
			}
		}
	}
	return result
}

func (matrix *Matrix) Inverse() Matrix {
	row, col := len(*matrix), len((*matrix)[0])

	dense := matrix.ConvertMatrixToDense()

	determinant := mat.Det(dense)
	inverseDeterminant := ModInverse(int(determinant), 26)
	if inverseDeterminant == -1 {
		return nil
	}

	var inverseDense mat.Dense
	err := inverseDense.Inverse(dense)
	if err != nil {
		return nil
	}

	inverseDense.Scale(determinant, &inverseDense)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			round := math.Round(inverseDense.At(i, j))
			inverseDense.Set(i, j, round)
		}
	}

	inverseDense.Scale(float64(inverseDeterminant), &inverseDense)

	var result Matrix
	result.ConvertDenseToMatrix(&inverseDense)

	return result
}

func (matrix *Matrix) ConvertMatrixToDense() *mat.Dense {
	row, col := len(*matrix), len((*matrix)[0])

	array := make([]float64, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			array[i*col+j] = float64((*matrix)[i][j])
		}
	}
	return mat.NewDense(row, col, array)
}

func (matrix *Matrix) ConvertDenseToMatrix(dense *mat.Dense) {
	row, col := dense.Dims()

	*matrix = InitializeMatrix(row, col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			(*matrix)[i][j] = int(dense.At(i, j))
		}
	}
}
