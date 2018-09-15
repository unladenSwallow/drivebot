package worldmap

import (
	"fmt"
)

const (
	xsize = 100
	ysize = 35
)

func PrintWorld() {
	fmt.Println("Printing The World....")
	// print the top
	printBand()
	// print the middle
	for y := 1; y < (ysize - 1); y++ {
		for x := 0; x < xsize; x++ {
			if x == 0 {
				fmt.Print("|")
			} else if x == (xsize - 1) {
				fmt.Println("|")
			} else {
				fmt.Print(" ")
			}
		}
	}
	// print the bottom
	printBand()
}

func printBand() {
	for t := 0; t < xsize; t++ {
		if t == (xsize - 1) {
			fmt.Println("-")
		} else {
			fmt.Print("-")
		}
	}
}
