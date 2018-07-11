package main

import (
	"flag"
	"fmt"
	"math"
)

/*
*    循环法 求x^n
 */
func power(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}

func main() {

	var init_free_cash_flow, increase_rate_of_free_cash_flow, discount_rate float64


	flag.Float64Var(&init_free_cash_flow, "init_free_cash_flow", 10.0, "init_free_cash_flow, e.g 10.0")
	flag.Float64Var(&increase_rate_of_free_cash_flow, "increase_rate_of_free_cash_flow", 0.15, "increase_rate_of_free_cash_flow")
	flag.Float64Var(&discount_rate, "discount_rate", 0.1, "discount_rate")
	flag.Parse()

	fmt.Println("========== param =============")
	/*
	year := 1988
	init_free_cash_flow = 8.28
	//i := 4 // 2021年的名义现金流
	i := 10
	increase_rate_of_free_cash_flow_arr := []float64{0.15}
	discount_rate_arr := []float64{0.09}
	forever_increase_rate_of_free_cash_flow := 0.05
	*/
	year := 2017
	init_free_cash_flow = 724
	i := 3 // 2021年的名义现金流
	increase_rate_of_free_cash_flow_arr := []float64{0.35, 0.3, 0.25, 0.2}
	discount_rate_arr := []float64{0.1, 0.12, 0.13}
	forever_increase_rate_of_free_cash_flow := 0.05

	fmt.Println("init_free_cash_flow:", init_free_cash_flow)
	fmt.Printf("start_year:%d\n", year)
	fmt.Printf("i:%d\n", i)


	/*
		(0) 2017: {free_cash_flow: 140.00} {after_discount: 140.00}
		(1) 2018: {free_cash_flow: 168.00} {after_discount: 152.73}
		(2) 2019: {free_cash_flow: 201.60} {after_discount: 166.61}
		(3) 2020: {free_cash_flow: 241.92} {after_discount: 181.76}
		(4) 2021: {free_cash_flow: 290.30} {after_discount: 198.28}
	*/


	for ii := 0; ii < len(increase_rate_of_free_cash_flow_arr); ii++ {
		increase_rate_of_free_cash_flow = increase_rate_of_free_cash_flow_arr[ii]
		fmt.Printf("========================R:%.0f%% ===============================\n", increase_rate_of_free_cash_flow*100)
		for jj := 0; jj < len(discount_rate_arr); jj++ {
			discount_rate = discount_rate_arr[jj]

			var q float64
			q = (1.0 + increase_rate_of_free_cash_flow) / (1.0 + discount_rate)

			if math.Abs(1.0-q) <= 0.0000000001 {
				fmt.Println("warning, divide zero, so stop!!!")
				return
			}

			discoun_q := (1.0 + increase_rate_of_free_cash_flow) / (1.0 + discount_rate)

			// 1. 算前面i年的现金流总量，并且贴现到起始年
			total_2021_discount := init_free_cash_flow * (1.0 - power(discoun_q, i)) / (1.0 - discoun_q)

			// 2. 算第i年，名义现金流, 以及当年的现金流贴现到起始年
			year_2021_free_cash_flow := init_free_cash_flow * power((1.0+increase_rate_of_free_cash_flow), i)
			after_discount := year_2021_free_cash_flow / (power((1.0 + discount_rate), i))
			fmt.Printf("********* R:%.0f%% D:%.0f%% A:%.0f%% **********\n", increase_rate_of_free_cash_flow*100,
				discount_rate*100, forever_increase_rate_of_free_cash_flow*100)
			fmt.Printf("%d: {this_year_free_cash_flow: %.2f} {this_year_after_discount: %.2f}\n", i+year, year_2021_free_cash_flow, after_discount)


			// 3. 算永续现金流
			forever_q := (1.0 + forever_increase_rate_of_free_cash_flow) / (1.0 + discount_rate)
			//这里往下多算了一年的永续增长率，然后再折现到今日
			forever_fcf := year_2021_free_cash_flow * (1+forever_increase_rate_of_free_cash_flow) / (1 - forever_q)
			total_all_discount_to_today := total_2021_discount + (forever_fcf / power((1+discount_rate), i+1))
			fmt.Printf("%d: total_fcf_before_and_include_year_%d_discount_to_%d: %.2f \n", i+year, i+year, year, total_2021_discount)
			fmt.Printf("%d: forever_fcf_discount_to_year_%d: %.2f , after_discount_to_%d: %.2f\n", i+year+1, i+year+1, forever_fcf, year, (forever_fcf / power((1+discount_rate), i+1)))
			fmt.Printf("%d: forever_fcf_discount_to_year_%d: %.2f , after_discount_to_%d: %.2f\n", i+year, i+year, forever_fcf/(1+discount_rate), year, (forever_fcf / power((1+discount_rate), i+1)))
			fmt.Printf("%d: CNY total_all_discount_to_%d: %.2f\n", year, year, total_all_discount_to_today)
			fmt.Printf("%d: USD total_all_discount_to_%d: %.2f\n", year, year, total_all_discount_to_today/6.5)
			fmt.Printf("%d: HKD total_all_discount_to_%d: %.2f\n", year, year, total_all_discount_to_today/0.83)
			fmt.Printf("%d: HKD per stock total_all_discount_to_%d: %.2f\n", year, year, total_all_discount_to_today/0.83/95.05)
			fmt.Print("\n")

		}

	}
}
