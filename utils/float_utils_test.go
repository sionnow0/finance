package utils_test

import "testing"
import (
	"../utils"
)

func Test_MySqrt(t *testing.T) {
	var num float64 = 16
	if utils.MySqrt(num) == 4 {
		t.Logf("succ")
	} else {
		t.Errorf("fail")
	}
}