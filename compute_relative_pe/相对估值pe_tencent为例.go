package main

import "fmt"

// TODO 搞个目录，然后加一个文件作为单元测试，以结果为准，然后可以大力地重构
func main() {
	// 715是2017年股东应占溢利
	net_profit_2017 := 715.1
	net_profit_2021 := 715.1 * 1.3* 1.3* 1.3 *1.3 // 这里算了成长4年，比dcf多算了一年
	r := 0.867 // 港币/人民币
	// 此刻20180712市值为3.61万亿港币，总股本95.05亿
	all_stock := 95.05 // (股票总数)
	// 静态pe是42.63
	now_pe := 42.63 //(富途上的pe为41.66，不对啊，算出来比市值少)
	a := net_profit_2017 * now_pe / r
	r_hkd_usd := 0.1274 // 港币兑美元=0.1274, 美元兑港币=7.848

	// ###########################

	fmt.Printf("2017 股东应占溢利:%.0f亿RMB {42pe: %.1fw HKD; %.0f亿USD; %.0f HKD}\n", net_profit_2017, a/10000, a*r_hkd_usd, a/all_stock)
	// 假设2021 pe介于25~30
	a_2021_25 := net_profit_2021 * 25 / r
	//fmt.Printf("2021 25pe: %.0fHKD\n", a_2021_25)
	a_2021_30 := net_profit_2021 * 30 / r
	//fmt.Printf("2021 30pe: %.0fHKD\n", a_2021_30)
	fmt.Printf("=============\n")
	fmt.Printf("net_profit_2021 := 715.1 * 1.3* 1.3* 1.3 *1.3 / (港币/人民币) = %.0fw亿港币; 25pe=%.0fw亿*25=%.0fw亿\n",
		net_profit_2021/r, net_profit_2021/r, net_profit_2021/r*25/10000)
	fmt.Printf("")
	fmt.Printf("2021 {25pe: %.1fw HKD ~ 30pe: %.1fw HKD}\n", a_2021_25/10000, a_2021_30/10000)
	fmt.Printf("2021 {25pe: %.0f USD ~ 30pe: %.0f USD}\n", a_2021_25*r_hkd_usd, a_2021_30*r_hkd_usd)
	fmt.Printf("2021 {25pe: %.0f HKD ~ 30pe: %.0f HKD}\n", a_2021_25/all_stock, a_2021_30/all_stock)
	fmt.Printf("2017 can buy price: {25pe:%.0f HKD ~ 30pe:%.0f HKD}\n", a_2021_25/all_stock/2, a_2021_30/all_stock/2)
	// 4年翻倍，72/4 = 18%
}
