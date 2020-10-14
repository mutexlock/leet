package problem0036

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValidSudokuRow(board, row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if !isValidSudokuCol(board, col) {
			return false
		}
	}

	for pod := 0; pod < 9; pod++ {
		if !isValidSudokuPod(board, pod) {
			return false
		}
	}
	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var nums [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			fmt.Println("Invalid row: ", row)
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var nums [10]bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			fmt.Println("Invalid col: ", col)
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuPod(board [][]byte, pod int) bool {
	var nums [10]bool

	row := (pod / 3) * 3
	col := (pod % 3) * 3

	for drow := 0; drow < 3; drow++ {
		for dcol := 0; dcol < 3; dcol++ {
			n := convertToNumber(board[row+drow][col+dcol])
			if n < 0 {
				continue
			}
			if nums[n] {
				fmt.Println("Invalid pod:", pod)
				fmt.Printf("Found duplicate number %d on row %d, col %d", n, row+drow, col+dcol)
				return false
			}
			nums[n] = true
		}
	}
	return true
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}


// 解法二 添加缓存，时间复杂度 O(n^2)
func isValidSudoku1(board [][]byte) bool {
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	// 遍历一次，添加缓存
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				num := board[r][c] - '0' - byte(1)
				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}
	return true
}