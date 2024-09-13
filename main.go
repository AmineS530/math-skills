package main

import (
	"os"

	skillhelper "math-skills/GoFiles"
)

var pData skillhelper.ProgData

func main() {
	if len(os.Args) != 2 {
		skillhelper.PrintErr("Program only accept one argument" +
			"\nEX: ./[program_name] <data.txt>")
	}
	pData.InitData()
	pData.DisplayResults()
}
