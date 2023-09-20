// imageManipulator/main.go
package main

import (
	"fmt"
	imagemod "myimageapp/imageManipulator"
)

func main() {
	// Create an ImageManipulator instance with an existing image
	im, err := imagemod.NewImageManipulatorWithImage("mustangs.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the modified image to a file
	im.SaveToFile("mustangs.png")
}
