package driver

import (
	"math/rand"
	"time"
)

const (
	Driver = "X"
)

var xpos, ypos int

func SetStartPosition(x, y int) {
	start := getRandInt()
	xpos = start % x
	ypos = start % y
}

func GetCurrentDriverPosition() (int, int) {
	return xpos, ypos
}

func MoveX(n int) {
	xpos += n
}

func MoveY(n int) {
	ypos += n
}

func RandomMove() {
	rand := getRandInt()
	if rand%8 == 0 {
		MoveX(-1)
		MoveY(-1)
	} else if rand%7 == 0 {
		MoveX(-1)
		MoveY(0)
	} else if rand%6 == 0 {
		MoveX(0)
		MoveY(-1)
	} else if rand%5 == 0 {
		MoveX(-1)
		MoveY(1)
	} else if rand%4 == 0 {
		MoveX(1)
		MoveY(-1)
	} else if rand%3 == 0 {
		MoveX(1)
		MoveY(1)
	} else if rand%2 == 0 {
		MoveX(0)
		MoveY(1)
	} else {
		MoveX(1)
		MoveY(0)
	}
}

func getRandInt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}
