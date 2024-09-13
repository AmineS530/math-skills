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

func (result ProgData) DisplayResults() {
	setNums(readFile(), &result)
	Calc(&result)
	fmt.Printf("Average: %d\n", int(math.Round(result.Res.Average)))
	fmt.Printf("Median: %d\n", int(math.Round(result.Res.Median)))
	fmt.Printf("Variance: %d\n", int(math.Round(result.Res.Variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(result.Res.Standard_Deviation)))
}

func PrintErr(str string) {
	fmt.Println("\033[1;38;5;221m╔════════════════════" + strings.Repeat("═", len(str)) + "════════════════════╗" + "\033[0m")
	fmt.Println("	    	\033[1;38;5;196mError:", str+"\033[0m")
	fmt.Println("\033[1;38;5;221m╚════════════════════" + strings.Repeat("═", len(str)) + "════════════════════╝" + "\033[0m")
	os.Exit(1)
}
