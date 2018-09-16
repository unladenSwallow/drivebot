package worldmap

import (
	"fmt"
	"math/rand"
	"time"

	"drivebot/1.0/driver"
)

const (
	xsize = 100
	ysize = 35
)

var board [][]string

func PrintWorld() {
	makeBoard()
	fmt.Println("Printing The World....")
	driverX, driverY := driver.GetCurrentDriverPosition()
	fmt.Println("", driverX, driverY)
	// print the middle
	for y := 0; y < ysize; y++ {
		for x := 0; x < xsize; x++ {
			if x == 0 {
				board[y][x] = "|"
				fmt.Print(board[y][x])
			} else if x == (xsize - 1) {
				board[y][x] = "|"
				fmt.Println(board[y][x])
			} else if y == 0 || y == ysize-1 {
				board[y][x] = "-"
				fmt.Print(board[y][x])
			} else if y == driverY && x == driverX {
				board[y][x] = driver.Driver
				fmt.Print(board[y][x])
			} else {
				if getRandomIntNow()%6 == 0 {
					wall := GetWall()
					board[y][x] = wall
					fmt.Print(board[y][x])
				} else {
					board[y][x] = " "
					fmt.Print(board[y][x])
				}
			}
		}
	}
}

func getRandomIntNow() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

func makeBoard() {
	board = make([][]string, ysize)
	for i := range board {
		board[i] = make([]string, xsize)
	}
}
