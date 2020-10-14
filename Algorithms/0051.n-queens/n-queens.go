package problem0051

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}

	cols := make([]bool, n)
	// 记录 '\' 方向的对角线的占用情况
	d1 := make([]bool, 2*n)
	// 记录 '/' 方向的对角线的占用情况
	d2 := make([]bool, 2*n)

	board := make([]string, n)

	res := [][]string{}

	dfs(0, cols, d1, d2, board, &res)

	return res
}

func dfs(r int, cols, d1, d2 []bool, board []string, res *[][]string) {

	if r == len(board) {
		tmp := make([]string, len(board))
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}

	n := len(board)

	for c := 0; c < len(board); c++ {
		// 把棋盘想象成
		//   以左上角为坐标原点 [0,0]
		//   C 轴正方向向右
		//   R 轴正方向向下
		// 的坐标系。
		// 这样的话，每个格子就都有了自己的坐标值 [c,r]
		//
		// 对于 '\' 方向的斜线而言
		//   同一个斜线上的格子，利用其坐标 [c,r] 计算 r-c 的结果相同，
		//   不同斜线上 r-c 的结果不同。
		//   所以可以用 r-c 代表不同的 '\' 方向的斜线的编号。
		//   但是 r-c 有可能是负值，无法作为切片的索引值
		//   所以 +n，可知 r-c+n >= 0
		//   所以，使用 r-c+n 作为 '\' 方向斜线切片的索引值。
		// 对于 '/' 方向的斜线而言
		//   同一个斜线上的格子，利用其坐标 [c,r] 计算 r+c 的结果相同，
		//   不同斜线上 r+c 的结果不同。
		//   所以可以用 r+c 代表不同的 '/' 方向的斜线的编号。
		//   所以，使用 r+c 作为 '/' 方向斜线切片的索引值。
		id1 := r - c + n
		id2 := r + c
		if !cols[c] && !d1[id1] && !d2[id2] {
			b := make([]byte, n)
			for i := range b {
				b[i] = '.'
			}
			b[c] = 'Q'
			board[r] = string(b)
			// 标记占用
			cols[c], d1[id1], d2[id2] = true, true, true

			dfs(r+1, cols, d1, d2, board, res)

			// 解除标记
			cols[c], d1[id1], d2[id2] = false, false, false
		}
	}
}

//是否可以在 board[row][col] 放置皇后
func isValid(board [][]byte,row,col int) bool {
	//同列是否有
	for i:=0; i< row; i++{
		if board[i][col] == 'Q'{
			return false
		}
	}
	//右上方
	for i,j :=row-1,col+1; i>=0 && j<len(board);{
		if board[i][j] == 'Q'{
			return false
		}
		i--
		j++
	}
	//左上方
	for i,j :=row-1,col-1; i>=0 && j>=0;{
		if board[i][j] == 'Q'{
			return false
		}
		i--
		j--
	}
	return  true
}

func solveNQueens1(n int) [][]string {
	var board [][]byte
	for i:=0;i<n;i++{
		b := make([]byte,n)
		for j:=0;j<n;j++{
			b[j] = '.'
		}
		board = append(board,b)
	}
	var res [][]string
	backtrack(board,0,n,&res)
	return res
}

func backtrack(board [][]byte,row int, n int,res *[][]string)  {
	if row == n{
		//添加结果
		*res = append(*res,toStrSlice(board))
		return
	}
	for col :=0;col <n; col++{
		board[row][col] = 'Q'
		if !isValid(board,row,col){
			continue
		}
		backtrack(board,row+1,n,res)
		board[row][col] = '.'

	}
}

func toStrSlice(board [][]byte) []string {
	var res []string
	for i:=0;i<len(board);i++{
		var str string
		for j:=0;j<len(board);j++ {
			str += string(board[i][j])
		}
		res = append(res,str)
	}
	return res
}