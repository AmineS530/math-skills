package skillhelper

import (
	"math"
	"sort"
)

// calls to all the calucaltion functions while passing a pointer to the struct
func Calc(pData *ProgData) {
	calcAvg(pData)
	calcMed(pData)
	calcVar(pData)
	calcStdDev(pData)
}

// calculates the average
func calcAvg(pData *ProgData) {
	sum := 0.0
	for _, num := range pData.Nums {
		sum += num
	}
	pData.Res.Average = sum / float64(len(pData.Nums))
}

// calculates the Median
func calcMed(pData *ProgData) {
	sort.Float64s(pData.Nums)
	nlen := len(pData.Nums)
	mid := nlen / 2
	if nlen%2 == 0 {
		pData.Res.Median = (pData.Nums[mid-1] + pData.Nums[mid]) / 2
		return
	}
	pData.Res.Median = pData.Nums[mid]
}

// calculates the Variance
func calcVar(pData *ProgData) {
	avg := pData.Res.Average
	var res float64
	for _, num := range pData.Nums {
		diff := num - avg
		res += diff * diff
	}
	pData.Res.Variance = res / float64(len(pData.Nums))
}

// calculates the Standard Deviation
func calcStdDev(pData *ProgData) {
	pData.Res.Standard_Deviation = math.Sqrt(pData.Res.Variance)
}
