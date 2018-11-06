package tmp_only_calculate

import (
	"testing"
	"math"
	"fmt"
)

func Test_test(t *testing.T) {

	v := math.Log(81) / math.Log(3)
	fmt.Printf("%.2f", v)
}

func Test_计算多少年收回成本(t *testing.T) {
	for i := 10; i < 300; i += 10 {
		s := float64(i)
		q := (100 + s) / 100.0
		a1 := 1.0
		res := 1.0 - (s*(1.0-q))/a1
		v := math.Log(res) / math.Log(q)
		fmt.Printf("pe:%d -- year:%.2f\n", i, v)
	}
	// ------------ pe为10，增长为10，7.27收回成本
	// ------------ pe为20，增长为20，8.83收回成本
	/*
		// 这个year计算的提前是，企业的增长率跟pe一致
	    pe:10 -- year:7.27
		pe:20 -- year:8.83
		pe:30 -- year:8.78
		pe:40 -- year:8.42
		pe:50 -- year:8.04
		pe:60 -- year:7.68
		pe:70 -- year:7.37
		pe:80 -- year:7.10
		pe:90 -- year:6.87
		pe:100 -- year:6.66
		pe:110 -- year:6.47
		pe:120 -- year:6.31
		pe:130 -- year:6.17
		pe:140 -- year:6.03
		pe:150 -- year:5.92
		pe:160 -- year:5.81
		pe:170 -- year:5.71
		pe:180 -- year:5.62
		pe:190 -- year:5.53
		pe:200 -- year:5.46
	 */
}


func Test_计算pe5(t *testing.T) {
		s := 5.0
		q := (100 + s) / 100.0
		a1 := 1.0
		res := 1.0 - (s*(1.0-q))/a1
		v := math.Log(res) / math.Log(q)
		fmt.Printf("%.2f\n", v)
	// ------------ pe为10，增长为10，7.27收回成本
	// ------------ pe为20，增长为20，8.83收回成本
}

