package main

import (
	"fmt"

	"./Colors"    // Importing Colors package
	"./GetPic"    // Importing GetPic package
	"./Grayscale" // Importing Grayscale package
	"./Text"      // Importing Text package
)

func main() {
	fmt.Println("Main function in main.go")

	Colors.ExampleFunction()
	Grayscale.AnotherFunction()
	Text.SomeType{}.SomeMethod()
	GetPic.DoSomething()
}
