package medium

import "math"

func getElementsSubmatrix(board [][]byte, row int, col int, size int) []byte {
	var result []byte
	for i := row; i < row+size; i++ {
		for j := col; j < col+size; j++ {
			result = append(result, board[i][j])
		}
	}
	return result
}

func getElementsRow(board [][]byte, row int) []byte {
	var result []byte
	for i := 0; i < 9; i++ {
		result = append(result, board[row][i])
	}
	return result
}

func getElementsCol(board [][]byte, col int) []byte {
	var result []byte
	for i := 0; i < 9; i++ {
		result = append(result, board[i][col])
	}
	return result
}

func verifyNonDuplicates(elements []byte) bool {
	exists := make(map[byte]bool)
	for _, value := range elements {
		if value == '.' {
			continue
		}
		if exists[value] {
			return false
		}
		exists[value] = true
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	size := 9
	subMatrixSize := int(math.Sqrt(float64(size)))

	// Check rows and columns
	for i := 0; i < size; i++ {
		if !verifyNonDuplicates(getElementsRow(board, i)) {
			return false
		}
		if !verifyNonDuplicates(getElementsCol(board, i)) {
			return false
		}
	}

	for i := 0; i < size; i += subMatrixSize {
		for j := 0; j < size; j += subMatrixSize {
			if !verifyNonDuplicates(getElementsSubmatrix(board, i, j, subMatrixSize)) {
				return false
			}
		}
	}

	return true
}
