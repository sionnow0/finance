package utils_test

import "testing"
import (
	"../utils"
	"math"
)

func Test_Power(t *testing.T) {
	// 2^4 = 16
	if utils.G_Accuracy.Equal(math.Pow(2, 4), 16) {
		t.Logf("succ")
	} else {
		t.Errorf("err")
	}
}

func Test_Power_v1(t *testing.T) {

	// 16^(1/4) = 2
	if utils.G_Accuracy.Equal(math.Pow(16, float64(1.0/4)), 2) {
		t.Logf("succ")
	} else {
		t.Errorf("err:%v", math.Pow(16, 1.0/4))
	}
}

func Test_Power_v2(t *testing.T) {
	// 16^(1/2) = 4 // 注意整数相除会变为整数
	if utils.G_Accuracy.Equal(math.Pow(16, 1.0/2), 4) {
		t.Logf("succ")
	} else {
		t.Errorf("err:%v", math.Pow(16, 1.0/2))
	}
}

func TestGetRatio(t *testing.T) {
	// 1.3 * 1.3 * 1.3 = 2.197, 增长率30%
	if utils.G_Accuracy.Equal(utils.GetRatio(1, 2.197, 3), 30) {
		t.Logf("succ")
	} else {
		t.Errorf("err")
	}
	t.Log(utils.GetRatio(1, 5, 10)) //u兄计算的那个表格，没有给增长率，所以我算了下，相当于17.5%
	t.Log(utils.GetRatio(1, 3, 10)) //10年3倍，相当于11.6%
	t.Log(utils.GetRatio(1, 2, 10)) //10年翻倍，相当于7.17%
	t.Log(utils.GetRatio(1, 2, 3)) //3年翻倍，相当于26%
}