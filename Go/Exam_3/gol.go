package main

import (
	"os"
	"strconv"
    "os/exec"
	"fmt"
	"io/ioutil"
	"time"
)

// var rowsT int
// var colsT int

func MakeBoard(rows int, cols int) (board [][]bool) {
	// colsT := len(board)
	// rowsT := len(board[1]) 
	grid2D := make([][]bool, cols)
    for i := 0; i < cols; i++ {
        grid2D[i] = make ([]bool, rows)
    }
	return grid2D
}

// go run inputs/01.txt 11 40 100 > out01.txt” and “diffout01.txt correct01.txt
// var xdim int = rows
// var ydim int = cols

// grid2D := make([][]int, xdim)
// for i := 0; i < xdim; i++ {
//     grid2D[i] = make ([]int, ydim)
// }


func BoardToString(board [][]bool) string {
	// var myStr string
	myStr := ""
	colsT := len(board)
	rowsT := len(board[1]) 
	// myStr := make([]string, 0, rows*cols)
	for i := 1; i<colsT-1;i++ {
		for j := 1; j<rowsT-1; j++ {
			if board[i][j] == true {
				myStr = myStr + "*"
				// myStr = append(myStr, "*")
				// fmt.Println("*")
			} else if board[i][j] == false {
				myStr = myStr + " "
				// myStr = append(myStr, " ")
				// fmt.Println(" ")
			}
		}
		myStr = myStr + "\n"
		// myStr = append(myStr, "\n")
		// fmt.Println("\n")
	}
	// fmt.Println("here")
	// fmt.Println(len(myStr))
	return myStr
}

// func BoardToStringParallel(board [][]bool) string {
// 		xdim := len(board)
// 		ydim := len(board[1])
	    
// 	    doneA := make(chan bool)
//         for i := ydim-1; i>=0 ; i-- {
//           go parallelStepTwo_v1(board, i, xdim, doneA)
//         }

//         for i := ydim-1; i>=0; i-- {
//           <-doneA  
//         }
//         doneB := make(chan bool)
//         for i := 0; i < xdim; i++ {
//           go parallelStepTwo_v2(board, i, ydim, doneB)
//         }
//         for i := 0; i < xdim; i++ {
//           <-doneB
//         }
// }

// func parallelStepTwo_v2(grid2D [][]int, temp int, ydim int, done chan bool){
//     for i := ydim-2; i>=0; i-- {
//         grid2D[temp][i] = grid2D[temp][i+1] + grid2D[temp][i]
//     }
//     done <- true
// }


// func parallelStepTwo_v1(grid2D [][]int, temp int, xdim int, done chan bool){
//     for i := 1; i<xdim;i++ {
//         grid2D[i][temp] = grid2D[i-1][temp] + grid2D[i][temp]
//     }
//     done <- true
// }


////////////////////////////
func StringToBoard(str []byte, board [][]bool) {
	// fmt.Println(len(board))
	// fmt.Println(len(board[1]))
	// fmt.Println(len(str))

	colsT := 1
	rowsT := 1

	for i:= 0; i < len(str)-1; i++ {

		if colsT >= len(board) {
			colsT = 1
		}
		if rowsT >= len(board[1]) {
			rowsT = 1
		}

		if str[i] == '1' {
			// fmt.Println("HERE")
			board[colsT][rowsT] = true
			rowsT++
		} else if (str[i] == ' ') {
			// fmt.Println(" HERE 2")
			// fmt.Println(colsT,rowsT)
			board[colsT][rowsT] = false
			rowsT++
		} else {
			// fmt.Println(" HERE 3")
			colsT++
			rowsT=1
		}
	}
}


// if (str i == 49)
// 	board[row][col]
// 	col++

// else if {str i == 32
// 	board
// col++
// } else {
	// row++
	// col = 1
// }

func NextCellState(board [][]bool, row int, col int) bool {
	// col := len(board)
	// row := len(board[1])
	// fmt.Println(len(board)) 
	// fmt.Println(len(board[1])) 
	var cellState bool
	// if row+1 >= len(board[1]) {
	// 	row = row-2
	// }
	// fmt.Println(col, row)
	if (row+1 < len(board[1])) {
		cellState = board[col][row]
		neighbours := 0
		if board[col-1][row-1] == true {
			neighbours++
		}
		if board[col][row-1] == true {
			neighbours++
		}
		if board[col+1][row-1] == true {
			neighbours++
		}
		if board[col-1][row] == true {
			neighbours++
		}
		if board[col+1][row] == true {
			neighbours++
		}
		// if board[col-1][row+1] == true {
		if board[col-1][row+1] == true {
			neighbours++
		}
		// if board[col][row+1] == true {
		if board[col][row+1] == true {
			neighbours++
		}
		// if board[col+1][row+1] == true {
		if board[col+1][row+1] == true {
			neighbours++
		}

		if cellState == true {
			if neighbours < 2 {
				cellState = false
				// board[row][col] = false
			} else if neighbours == 2 || neighbours == 3 {
				cellState = true
				// board[row][col] = true
			} else if neighbours >= 3 {
				cellState = false
				// board[row][col] = false
			}
		} else {
			if neighbours == 3 {
				cellState = true
				// board[row][col] = true
			}
		}
	}
	return cellState
}

func NextGameState(oldBoard [][]bool, newBoard [][]bool) {
	colsT := len(oldBoard)
	rowsT := len(oldBoard[1])
	for i := 1; i<colsT-1;i++ {
		for j := 1; j<rowsT-1; j++ {
			newBoard[i][j] = NextCellState(oldBoard, i, j)
		}
	}
}


 // You will use this main after finishing part 5. 
func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run gol.go <filename> <rows> <cols> <iterations>")
		return
	}
	rows, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
    cols, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
    iters, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println(err)
		return
	}
	// rowsT = rows
	// colsT = cols
	gofile, _ := ioutil.ReadFile(os.Args[1])
	oldBoard := MakeBoard(rows, cols)
	StringToBoard(gofile, oldBoard)
	newBoard := MakeBoard(rows, cols)
	for i := 0; i < iters; i++ {
		NextGameState(oldBoard, newBoard)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
        cmd.Run()
		fmt.Print(BoardToString(newBoard))
		oldBoard, newBoard = newBoard, oldBoard
		time.Sleep(time.Second / 30)
	}
}


// func StringToBoard_2(str []byte, board [][]bool) {
// 	// colsT := len(board)
// 	// rowsT := len(board[1]) 
// 	j := 1
// 	for i := 1; i<len(str);i++ {
// 		// for j := 1; j<cols-1; j++ {
// 			if str[i] == '\n' {
// 				j++
// 				i = 1
// 			} else if str[i] == '1' {
// 				board[i][j] = true
// 			} else if str[i] == ' ' {
// 				board[i][j] = false
// 			}
// 		// }
// 	}
// }


// func StringToBoard_2(str []byte, board [][]bool) {
// 	// colsT := len(board)
// 	// rowsT := len(board[1]) 
// 	j := 1
// 	i := 1
// 	for k := 1; k<len(str);k++ {
// 		// for j := 1; j<cols-1; j++ {
// 			// if str[i] == '\n' {
// 			// 	j++
// 			// 	i = 1
// 			// } else if str[i] == '1' {
// 			if str[i] == '*' {
// 				board[i][j] = true
// 				i++
// 			} else if str[i] == ' ' {
// 				board[i][j] = false
// 				i++
// 			} else {
// 				j++
// 				i = 1
// 			}
// 		// }
// 	}
// }

