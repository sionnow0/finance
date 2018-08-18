package utils

import "math"

// 求x的开方，16结果为4
func MySqrt(x float64) float64 {
	z := float64(1)
	tmp := float64(0)
	for math.Abs(tmp - z) > 0.0000000001 {
		tmp = z
		z = (z + x/z)/2
	}
	return z
}

/*
*    循环法 求x^n
 */
func Power(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}

// 给出起始金额，终止金额，年数，求出这些年的增长率,
// 增长率已经算好百分比，比如10%，就返回10
func GetRatio(start_num, end_num, year_num float64) float64 {
	return (math.Pow(end_num/start_num, 1.0/year_num) - 1.0) * 100
}