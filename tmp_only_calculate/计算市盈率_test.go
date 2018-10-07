package tmp_only_calculate

import (
	"testing"
	"fmt"
)

func Test_静态pe(t *testing.T) {
	hk_ratio := 0.8797
	pe := 305.0 / (7.499 / hk_ratio)
	fmt.Println(pe)
	pe = 305.0 / (7.598 / hk_ratio)
	fmt.Println(pe)
	profit_2017 := 715.10

	yi := 100000000.0 // 1亿
	pe = 305 * 95.21 * yi / (profit_2017 * yi / hk_ratio)
	fmt.Println(pe)
	/*
	35.77923723163089
	35.3130429060279
	35.723188763809254
	富途怎么静态pe是33.43, 富途傻逼乱算
	*/
}

func Test_TTM_pe(t *testing.T) {
	hk_ratio := 0.8797
	var pe float64
	profit_2018_q2q1 := 411.57
	profit_2017_q4q3 := 180 + 207.9
	profit := profit_2017_q4q3 + profit_2018_q2q1
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = 305 * 95.21 * yi / (profit * yi / hk_ratio)
	fmt.Println(pe)
	/*
	31.95323437402279
	富途怎么静态pe是29.64, 富途傻逼乱算
	*/
}

func Test_wangyi_static_pe(t *testing.T) {
	var pe float64
	//profit_2017 := 16.41 // 美元
	dollor_ratio := 6.8688
	profit_2017 := 107.08  / dollor_ratio
	profit := profit_2017
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = 283.026 * yi / (profit * yi)
	fmt.Println(pe)
	/*
	17.24
	富途怎么静态pe是17.36, 富途傻逼乱算
	*/
}

func Test_wangyi_ttm_pe(t *testing.T) {
	var pe float64
	dollor_ratio := 6.8688
	profit_2018_q2q1 := 21.07 + 7.52  // 人民币
	profit_2017_q4q3 := 12.86 + 25.27 // 人民币
	profit := (profit_2017_q4q3 + profit_2018_q2q1) / dollor_ratio

	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = 283.026 * yi / (profit * yi)
	fmt.Println(pe)
	/*
		profit: 9.71
		29.13
	富途怎么ttm pe是23.726
	*/
}

func Test_google_static_pe(t *testing.T) {
	var pe float64
	profit_2017 := 126.62
	marker_value := 8051.42
	profit := profit_2017
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = marker_value * yi / (profit * yi)
	fmt.Printf("%.2f\n", pe)
	/*
	63.59
	富途64.297
	*/
}

func Test_google_ttm_pe(t *testing.T) {
	var pe float64
	profit_2018_q2q1 := 31.95 + 94.01
	profit_2017_q4q3 := -30.2 + 67.32
	profit := (profit_2017_q4q3 + profit_2018_q2q1)
	marker_value := 8051.42
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = marker_value * yi / (profit * yi)
	fmt.Printf("%.2f\n", pe)
	/*
		profit: 163.07
		49.37
	富途怎么ttm pe是49.993
	*/
}

func Test_apple_ttm_pe(t *testing.T) {
	var pe float64
	profit_2018_q2q1 := 115.19 + 138.22
	profit_2017_q4q3 := 200.65 + 107.14
	profit := (profit_2017_q4q3 + profit_2018_q2q1)
	marker_value := 10833.0
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = marker_value * yi / (profit * yi)
	fmt.Printf("%.2f\n", pe)
	/*
		profit: 561.2
		19.30
	富途怎么ttm pe是20.279
	*/
}

func Test_fb_ttm_pe(t *testing.T) {
	var pe float64
	profit_2018_q2q1 := 51.06 + 49.88
	profit_2017_q4q3 := 42.69 + 47.07
	profit := (profit_2017_q4q3 + profit_2018_q2q1)
	marker_value := 4542.46
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = marker_value * yi / (profit * yi)
	fmt.Printf("%.2f\n", pe)
	/*
		profit: 190.7
		23.82
	富途怎么ttm pe是24.354
	*/
}

func Test_amazon_ttm_pe(t *testing.T) {
	var pe float64
	profit_2018_q2q1 := 25.34 + 16.29
	profit_2017_q4q3 := 18.57 + 2.56
	profit := (profit_2017_q4q3 + profit_2018_q2q1)
	marker_value := 9216.60
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = marker_value * yi / (profit * yi)
	fmt.Printf("%.2f\n", pe)
	/*
		profit: 62.76
		146.85
	富途怎么ttm pe是147.628
	*/
}

func Test_kele_ttm_pe(t *testing.T) {
	var pe float64
	profit_2018_q2q1 := 23.16 + 13.68
	profit_2017_q4q3 := -27.52 + 14
	profit := (profit_2017_q4q3 + profit_2018_q2q1)
	marker_value := 1951.24
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = marker_value * yi / (profit * yi)
	fmt.Printf("%.2f\n", pe)
	/*
		profit: 23.320000000000004
		83.67
	*/
}

func Test_baba_ttm_pe(t *testing.T) {
	var pe float64
	dollor_ratio := 6.8688
	profit_2018_q2q1 := 86.85 + 75.61
	profit_2017_q4q3 := 240.73 + 176.68
	profit := (profit_2017_q4q3 + profit_2018_q2q1) / dollor_ratio
	marker_value := 4008.29
	fmt.Println("profit:", profit)
	yi := 100000000.0 // 1亿
	pe = marker_value * yi / (profit * yi)
	fmt.Printf("%.2f\n", pe)
	/*
		profit: 84.42
		47.48
	富途怎么ttm pe是46.199
	*/
}