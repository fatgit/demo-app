package main

import (
	"fmt"
	"os"
	"time"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("ofominykh is Awesome!!!", "larry3d", "yellow", true)
	fmt.Println("Checking cheking")

	if secret := os.Getenv("SECRET"); secret != "" {
		myFigure = figure.NewColorFigure(fmt.Sprintf("Secret value is: %s", secret), "larry3d", "yellow", true)
	}
	myFigure.Print()
	time.Sleep(10 * time.Hour)
}
