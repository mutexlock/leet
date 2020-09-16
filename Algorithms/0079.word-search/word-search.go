package problem0079
//
//func exist(board [][]byte, word string) bool {
//	m := len(board)
//	if m == 0 {
//		return false
//	}
//
//	n := len(board[0])
//	if n == 0 {
//		return false
//	}
//
//	if len(word) == 0 {
//		return false
//	}
//
//	var dfs func(int, int, int) bool
//	dfs = func(r, c, index int) bool {
//		if len(word) == index {
//			return true
//		}
//
//		if r < 0 || m <= r || c < 0 || n <= c || board[r][c] != word[index] {
//			return false
//		}
//
//		temp := board[r][c]
//		board[r][c] = 0
//
//		if dfs(r-1, c, index+1) ||
//			dfs(r+1, c, index+1) ||
//			dfs(r, c-1, index+1) ||
//			dfs(r, c+1, index+1) {
//			return true
//		}
//
//		board[r][c] = temp
//
//		return false
//	}
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if dfs(i, j, 0) {
//				return true
//			}
//		}
//	}
//
//	return false
//}


func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0{
		return false
	}
	column := len(board[0])
	if column == 0{
		return false
	}
	for i:=0;i< len(board);i++{
		for j:=0;j < column;j++{
			exist := dfs(board,i,j,word,0)
			if exist {
				return true
			}
		}
	}
	return false
}

func dfs(board[][]byte, i,j int,word string, index int) bool{
	if len(word) == index{
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	exist := false
	if board[i][j] == word[index]{
		board[i][j] = 0
		if dfs(board,i-1,j, word,index+1) ||  dfs(board,i,j-1, word,index+1) ||  dfs(board,i+1,j, word,index+1) ||  dfs(board,i,j+1, word,index+1){
			exist = true
		}
		board[i][j] = word[index]
	}
	return exist
}