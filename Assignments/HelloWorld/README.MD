# Notes for GoLang



A brief description of your project goes here.

## Table of Contents for Notes

- [Setting Up](#Setting-Up )
  - [Installation](#installation)
  - [Usage](#usage)
- [Resources](#Resources)


## Setting Up 
Note: This is with <b>VS Code</b> 

Setting up GoLang we need to go to <a>https://go.dev/dl/</a> and download 
the following package for the according operating system you are using. 
### Installation
- Step1: After doing so we will click next and let the package install
- Step2: Next download the Go package extension for VSCode 
<img src="/images/go-extpic.png"width=100px>

- Step3: Create a folder and file <b>Example File Name</b> filename.go
- Step4: Then we need to control+shift+p type <b>go update</b> check mark all 
and download all updates. 

### Usage

- Starting a project is the same as other unless that has to do with multiple repos.
- Assuming the process is basic we will create a folder like so <b>/folderNameHere/fileNameHere.go</b>. Ending it with go. 

- Then will create small helloWorld to show a proper skeleton

```
// library for starting program
package main

//provides functions for formatting and printing
//values. Examples: String formatting Sprintf()=formats a string
//Errorf()= create formatted errors, Scanning Scan= read and parse input
// from the console, Fprintf()= write formatted output like a file or network
// os.Stdin(), os.Stdout(), and os.Stderr().
import "fmt"

//function main equivalent to int main.
func main() {
	//Prints a line saying Hello World!
	fmt.Println("Hello World")
}

//to run the program we write in terminal the following:
// go run nameoffile.go

```

## Resources

- Go Playground: https://go.dev/play/ 






