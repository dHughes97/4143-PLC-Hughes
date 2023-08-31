// library for starting program
// package main

// //provides functions for formatting and printing
// //values. Examples: String formatting Sprintf()=formats a string
// //Errorf()= create formatted errors, Scanning Scan= read and parse input
// // from the console, Fprintf()= write formatted output like a file or network
// // os.Stdin(), os.Stdout(), and os.Stderr().
// import "fmt"

// //function main equivalent to int main.
// func main() {
// 	//Prints a line saying Hello World!
// 	fmt.Println("Hello World")
// }

//to run the program we write in terminal the following:
// go run nameoffile.go

package main

import "fmt"

//creating a struct of a person with variables name and age
type Person struct {
	Name string
	Age  int
}

//Creating a var p that is a Person, Making a function called introduce(),
// which will print a line importing the correct data types using concatenation
func (p Person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

//creating instances of a person and introducing them
func main() {
	person := Person{Name: "Bob", Age: 25}
	person.introduce()
}
