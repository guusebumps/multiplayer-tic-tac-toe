package main

import (
	"fmt"
)

var board = [3][3]string{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

var currentPlayer = "X"

func main() {
	for {
		printBoard()
		var row, col int
		fmt.Printf("Player %s, enter row and column: ", currentPlayer)
		fmt.Scan(&row, &col)
		if board[row][col] == "-" {
			board[row][col] = currentPlayer
			if checkWinner() {
				printBoard()
				fmt.Printf("Player %s wins!\n", currentPlayer)
				break
			}
			switchPlayer()
		}
	}
}

func printBoard() {
	for _, row := range board {
		fmt.Println(row)
	}
}

func checkWinner() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}
	return false
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}
