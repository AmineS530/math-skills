package skillhelper

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ProgData struct {
	Nums []float64
	Res  calculations
}

type calculations struct {
	Average, Median, Variance, Standard_Deviation float64
}

// reads the the input file from os.args
func readFile() []string {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		PrintErr(err.Error())
	}
	if len(data) == 0 {
		PrintErr("Error: The File is Empty")
	}
	processedData := strings.ReplaceAll(string(data), " ", "")
	return strings.Split(processedData, "\n")
}

// goes through the read input file and convert every number from string to float64
// and stores them in arrat of floats located in the struct used
func setNums(FileData []string, pData *ProgData) {
	for i := 0; i < len(FileData); i++ {
		if FileData[i] != "" && i < len(FileData) {
			num, err := strconv.ParseFloat(FileData[i], 64)
			if err != nil {
				PrintErr(err.Error())
			}
			pData.Nums = append(pData.Nums, num)
		}
	}
}

// calls on functions to get the appropriate results and prints them after rounding them
func (result ProgData) DisplayResults() {
	setNums(readFile(), &result)
	Calc(&result)
	fmt.Printf("Average: %d\n", int(math.Round(result.Res.Average)))
	fmt.Printf("Median: %d\n", int(math.Round(result.Res.Median)))
	fmt.Printf("Variance: %d\n", int(math.Round(result.Res.Variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(result.Res.Standard_Deviation)))
}

// prints errors in a styled way
func PrintErr(str string) {
	fmt.Println("\033[1;38;5;221m╔════════════════════" + strings.Repeat("═", len(str)) + "════════════════════╗" + "\033[0m")
	fmt.Println("	    	\033[1;38;5;196mError:", str+"\033[0m")
	fmt.Println("\033[1;38;5;221m╚════════════════════" + strings.Repeat("═", len(str)) + "════════════════════╝" + "\033[0m")
	os.Exit(1)
}
