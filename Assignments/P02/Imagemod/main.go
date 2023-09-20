
package main

import (
	"fmt"
	"myimageapp/imagemod"
	"github.com/T-Gilbert/4143-PLC-Gilbert/tree/main/Assignments/P02/Imagemod"
)

func main() {
	im, err := imagemod.NewImageManipulatorWithImage(LazyBear.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	im.DrawRectangle(150, 50, 560, 411)

	im.SaveToFile("LazyBear.png")
}
