package main

import "fmt"

func main() {
	/*
	a := 3365.12 / 2626.16
	fmt.Println(a)
	b := 1061 - 964 + 266
	fmt.Printf("%d\n", b)


	a = 6000.0/4705 * 386
	*/
	// 715是2017年股东应占溢利
	net_profit_2017 := 715.1
	net_profit_2021 := 715.1 * 1.3* 1.3* 1.3 * 1.3
	r := 0.84 // 港币/人民币
	// 此刻20180712市值为3.61万亿港币，总股本95.05亿
	all_stock := 95.05
	// 静态pe是42.63
	j_pe := 42.63 //(富途上的pe为41.66，不对啊，算出来比市值少)
	a := net_profit_2017 * j_pe / r
	r_hkd_usd := 0.1274 // 港币兑美元=0.1274, 美元兑港币=7.848
	fmt.Printf("2017 36pe: %.0fHKD %.0fUSD\n", a, a*r_hkd_usd)
	// 假设2021 pe介于25~30
	a_2021_25 := net_profit_2021 * 25 / r
	//fmt.Printf("2021 25pe: %.0fHKD\n", a_2021_25)
	a_2021_30 := net_profit_2021 * 30 / r
	//fmt.Printf("2021 30pe: %.0fHKD\n", a_2021_30)
	fmt.Printf("2021 {25pe: %.0fHKD ~ 30pe: %.0fHKD}\n", a_2021_25, a_2021_30)
	fmt.Printf("2021 {25pe: %.0fUSD ~ 30pe: %.0fUSD}\n", a_2021_25*r_hkd_usd, a_2021_30*r_hkd_usd)
	fmt.Printf("2021 {25pe: %.0fHKD ~ 30pe: %.0fHKD}\n", a_2021_25/all_stock, a_2021_30/all_stock)
	fmt.Printf("2017 can buy price: {25pe:%.0fHKD ~ 30pe:%.0fHKD}\n", a_2021_25/all_stock/2, a_2021_30/all_stock/2)
	// 4年翻倍，72/4 = 18%



	}
