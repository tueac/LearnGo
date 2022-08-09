package main

// 二维矩阵
func MatrixAdd(A [8][5]int, B [8][5]int) [8][5]int {
	C := [8][5]int{}
	for row := 0; row < 8; row++ {
		for col := 0; col < 5; col++ {
			C[row][col] = A[row][col] + B[row][col]
		}
	}
	return C
}
