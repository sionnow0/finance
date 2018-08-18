package main_test

import "testing"
import (
	"math"
	"../utils"
)

var G_All_Stock_Num float64 = 95.05 // (20180818股票总数)
var G_My_Costing float64 = 336 // (成本)
// 四舍五入
func round(num float64) uint32 {
	return uint32(math.Floor(num + 0.5))
}

func compute_to_2021(ratio float64) float64 {
	// 715是2017年股东应占溢利
	// 注意2017有200亿左右来自投资收益，以后可否持续呢？
	net_profit_2017 := 715.1
	net_profit_2021 := net_profit_2017 * ratio * ratio * ratio * ratio // 这里算了成长4年，比dcf多算了一年
	return net_profit_2021
}

func Test_pe_v1(t *testing.T) {
	// ###########################

	// 715是2017年股东应占溢利
	net_profit_2021 := compute_to_2021(1.3) // 这里算了成长4年，比dcf多算了一年
	r := 0.867      // 港币/人民币
	// 此刻20180712市值为3.61万亿港币，总股本95.05亿
	var pe_range_l, pe_range_r float64 = 20, 30

	// ###########################
	a_2021_25 := net_profit_2021 * pe_range_l / r
	//fmt.Printf("2021 25pe: %.0fHKD\n", a_2021_25)
	a_2021_30 := net_profit_2021 * pe_range_r / r

	hard_code := uint32(248)
	if round(a_2021_25/G_All_Stock_Num/2) == hard_code {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_l), hard_code*2, hard_code)
	} else {
		t.Errorf("2017 can buy price { %dpe: 248 HKD, but:%d }", round(pe_range_l), round(a_2021_25/G_All_Stock_Num/2))
	}
	hard_code1 := uint32(372)
	if round(a_2021_30/G_All_Stock_Num/2) == hard_code1 {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_r), hard_code1*2, hard_code1)
	} else {
		t.Errorf("2017 can buy price { %dpe: 372 HKD, but:%d }", round(pe_range_r), round(a_2021_30/G_All_Stock_Num/2))
	}

	year_num := 4.0
	t.Logf("公司30%%增长，收益增长率范围为%.1f%% ~ %.1f%%",
		utils.GetRatio(G_My_Costing, 496, year_num),
		utils.GetRatio(G_My_Costing, 744, year_num))
}

func Test_pe_v2(t *testing.T) {
	// ###########################

	// 715是2017年股东应占溢利
	net_profit_2021 := compute_to_2021(1.3) // 这里算了成长4年，比dcf多算了一年
	r := 0.867      // 港币/人民币
	// 此刻20180712市值为3.61万亿港币，总股本95.05亿
	var pe_range_l, pe_range_r float64 = 25, 30

	// ###########################
	a_2021_25 := net_profit_2021 * pe_range_l / r
	//fmt.Printf("2021 25pe: %.0fHKD\n", a_2021_25)
	a_2021_30 := net_profit_2021 * pe_range_r / r

	hard_code := uint32(310)
	if round(a_2021_25/G_All_Stock_Num/2) == hard_code {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_l), hard_code*2, hard_code)
	} else {
		t.Errorf("2017 can buy price { %dpe: 248 HKD, but:%d }", round(pe_range_l), round(a_2021_25/G_All_Stock_Num/2))
	}
	hard_code1 := uint32(372)
	if round(a_2021_30/G_All_Stock_Num/2) == hard_code1 {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_r), hard_code1*2, hard_code1)
	} else {
		t.Errorf("2017 can buy price { %dpe: 372 HKD, but:%d }", round(pe_range_r), round(a_2021_30/G_All_Stock_Num/2))
	}

	year_num := 4.0
	t.Logf("公司30%%增长，收益增长率范围为%.1f%% ~ %.1f%%",
		utils.GetRatio(G_My_Costing, 620, year_num),
		utils.GetRatio(G_My_Costing, 744, year_num))
}

func Test_pe_v3(t *testing.T) {

	// ###########################
	// 715是2017年股东应占溢利
	net_profit_2021 := compute_to_2021(1.2) // 这里算了成长4年，比dcf多算了一年
	r := 0.867      // 港币/人民币
	// 此刻20180712市值为3.61万亿港币，总股本95.05亿
	var pe_range_l, pe_range_r float64 = 20, 30
	// ###########################

	a_2021_25 := net_profit_2021 * pe_range_l / r

	a_2021_30 := net_profit_2021 * pe_range_r / r

	hard_code := uint32(180)
	if round(a_2021_25/G_All_Stock_Num/2) == hard_code {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_l), hard_code*2, hard_code)
	} else {
		t.Errorf("2017 can buy price { %dpe: 248 HKD, but:%d }", round(pe_range_l), round(a_2021_25/G_All_Stock_Num/2))
	}
	hard_code1 := uint32(270)
	if round(a_2021_30/G_All_Stock_Num/2) == hard_code1 {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_r), hard_code1*2, hard_code1)
	} else {
		t.Errorf("2017 can buy price { %dpe: 372 HKD, but:%d }", round(pe_range_r), round(a_2021_30/G_All_Stock_Num/2))
	}

	year_num := 4.0
	t.Logf("公司20%%增长，收益增长率范围为%.1f%% ~ %.1f%%",
		utils.GetRatio(G_My_Costing, 360, year_num),
		utils.GetRatio(G_My_Costing, 540, year_num))
}

func Test_pe_v4(t *testing.T) {

	// ###########################
	// 715是2017年股东应占溢利
	net_profit_2021 := compute_to_2021(1.2) // 这里算了成长4年，比dcf多算了一年
	r := 0.867      // 港币/人民币
	// 此刻20180712市值为3.61万亿港币，总股本95.05亿
	var pe_range_l, pe_range_r float64 = 25, 30
	// ###########################

	a_2021_25 := net_profit_2021 * pe_range_l / r

	a_2021_30 := net_profit_2021 * pe_range_r / r

	hard_code := uint32(225)
	if round(a_2021_25/G_All_Stock_Num/2) == hard_code {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_l), hard_code*2, hard_code)
	} else {
		t.Errorf("2017 can buy price { %dpe: 248 HKD, but:%d }", round(pe_range_l), round(a_2021_25/G_All_Stock_Num/2))
	}
	hard_code1 := uint32(270)
	if round(a_2021_30/G_All_Stock_Num/2) == hard_code1 {
		t.Logf("2017 can buy price { %dpe: %d/%d HKD }", round(pe_range_r), hard_code1*2, hard_code1)
	} else {
		t.Errorf("2017 can buy price { %dpe: 372 HKD, but:%d }", round(pe_range_r), round(a_2021_30/G_All_Stock_Num/2))
	}

	year_num := 4.0
	t.Logf("公司20%%增长，收益增长率范围为%.1f%% ~ %.1f%%",
		utils.GetRatio(G_My_Costing, 450, year_num),
		utils.GetRatio(G_My_Costing, 540, year_num))
}



func _Test_pe(t *testing.T) {
	// 715是2017年股东应占溢利
	net_profit_2017 := 715.1
	net_profit_2021 := net_profit_2017 * 1.3 * 1.3 * 1.3 * 1.3 // 这里算了成长4年，比dcf多算了一年
	r := 0.867                                                 // 港币/人民币
	// 此刻20180712市值为3.61万亿港币，总股本95.05亿

	// 静态pe是42.63
	//now_pe := 42.63 //(富途上的pe为41.66，不对啊，算出来比市值少)
	//a := net_profit_2017 * now_pe / r
	//r_hkd_usd := 0.1274 // 港币兑美元=0.1274, 美元兑港币=7.848

	// ###########################
	a_2021_25 := net_profit_2021 * 25 / r
	//fmt.Printf("2021 25pe: %.0fHKD\n", a_2021_25)
	a_2021_30 := net_profit_2021 * 30 / r

	if round(a_2021_25/G_All_Stock_Num/2) == 310 {
		t.Log("2017 can buy price: 25pe:310 HKD")
	} else {
		t.Error("2017 can buy price: 25pe:310 HKD, but:%d", round(a_2021_25/G_All_Stock_Num/2))
	}
	if round(a_2021_30/G_All_Stock_Num/2) == 372 {
		t.Log("2017 can buy price: 30pe:372 HKD")
	} else {
		t.Error("2017 can buy price: 30pe:372 HKD, but:%d", round(a_2021_30/G_All_Stock_Num/2))
	}
}
