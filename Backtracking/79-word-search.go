/*

79. Word Search
Medium
Topics
Companies
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCCED'
Output: true

Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'SEE'
Output: true


Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCB'
Output: false
*/

package BackTracking

import "fmt"

var dir = [][]int {
	 {-1, 0},
	 {0, -1},
	 {1, 0},
	 {0, 1},
}

func search(board [][]byte, visited [][]bool, word string, index, x, y int) bool{

	if index == len(word)-1 {
		return board[x][y] == word[index]
	}

	if board[x][y] == word[index] {
		visited[x][y] = true
		for i:=0; i<4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && search(board, visited, word, index+1, nx, ny){
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x>=0 && y>=0 && x<len(board) && y<len(board[0])
}

func wordSearch(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i:=0; i<len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i, v := range board {
		for j := range v {
			if search(board, visited, word, 0, i, j){
				return true
			}
		}
	}
	return false
}

func Test() {
	board_01 := [][]byte{ {'A','B','C','E'}, {'S','F','C','S'}, {'A','D','E','E'}}

	fmt.Println(wordSearch(board_01, "ABCCED"))
}