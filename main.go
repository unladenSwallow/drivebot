package main

import (
	"fmt"

	"drivebot/1.0/driver"
	wm "drivebot/1.0/worldmap"
)

func main() {
	driver.SetStartPosition(100, 35)
	fmt.Println("Hello from drivebot 1.0!")
	wm.PrintWorld()
}
