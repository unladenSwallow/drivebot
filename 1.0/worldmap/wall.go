package worldmap

import (
	"math/rand"
	"time"
)

func GetWall() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Int()%3 == 0 {
		return "_"
	} else {
		return "|"
	}
}
