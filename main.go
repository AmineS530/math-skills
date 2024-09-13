package main

import (
	"os"

	skillhelper "math-skills/GoFiles"
)

// declaring the struct ill be using throught the program
var pData skillhelper.ProgData

// main :)
func main() {
	if len(os.Args) != 2 {
		skillhelper.PrintErr("Program only accept one argument\nEX: ./[program_name] <file_name.txt>")
	}
	// call the method that does all the work within it
	pData.DisplayResults()
}
