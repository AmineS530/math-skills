package skillhelper

import (
	"math"
	"sort"
)

func Calc(pData *ProgData) {
	calcAvg(pData)
	calcMed(pData)
	calcVar(pData)
	calcStdDev(pData)
}

func calcAvg(pData *ProgData) {
	sum := 0.0
	for _, num := range pData.Nums {
		sum += num
	}
	pData.Res.Average = sum / float64(len(pData.Nums))
}

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

func calcVar(pData *ProgData) {
	avg := pData.Res.Average
	var res float64
	for _, num := range pData.Nums {
		diff := num - avg
		res += diff * diff
	}
	pData.Res.Variance = res / float64(len(pData.Nums))
}

func calcStdDev(pData *ProgData) {
	pData.Res.Standard_Deviation = math.Sqrt(pData.Res.Variance)
}
