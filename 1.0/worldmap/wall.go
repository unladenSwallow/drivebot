package worldmap

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintWall() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Int()%2 == 0 {
		fmt.Print("_")
	} else {
		fmt.Print("|")
	}
}
