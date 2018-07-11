package main

import (
	"fmt"
)

type PeInfo struct {
	dw_year int
	now_pe  float64
	now_eps float64
}

func main() {

	//假定初始price为1CNY
	var dw_init_pe float64 = 37
	var growth_rate float64 = 35//百分比
	var year_num int = 5
	year_num++
	var stPeInfo [1024]PeInfo
	stPeInfo[0].dw_year = 0
	stPeInfo[0].now_eps = 1
	stPeInfo[0].now_pe = dw_init_pe

	for i := 1; i < year_num; i++ {
		stPeInfo[i].dw_year = i
		stPeInfo[i].now_eps = stPeInfo[i-1].now_eps * (100.0 + growth_rate) / 100.0
		stPeInfo[i].now_pe = dw_init_pe / stPeInfo[i].now_eps
	}

	fmt.Printf("year\t\t\t")
	for i := 0; i < year_num; i++ {
		//fmt.Printf("第%d年\t", i);
		fmt.Printf("%d\t\t", i)
	}
	fmt.Printf("\n")
	fmt.Printf("1CNY start\t\t")
	for i := 0; i < year_num; i++ {
		fmt.Printf("%.2f\t", stPeInfo[i].now_eps)
	}
	fmt.Printf("\n")
	fmt.Printf("pe \t\t\t\t")
	for i := 0; i < year_num; i++ {
		fmt.Printf("%.2f\t", stPeInfo[i].now_pe)
	}
	fmt.Printf("\n\n")

	for j := dw_init_pe; j >= 10; j -= 3 {
		fmt.Printf("priceOfpe%.0f \t", j)
		for i := 0; i < year_num; i++ {
			fmt.Printf("%.1f\t", stPeInfo[i].now_eps*j)
		}
		fmt.Printf("\n")
		fmt.Printf("rateOfgrow \t\t")
		for i := 0; i < year_num; i++ {
			tmp := stPeInfo[i].now_eps * j / (1*dw_init_pe) * 100.0
			fmt.Printf("%.1f%%\t", tmp)
		}
		fmt.Printf("\n\n")
	}
}
