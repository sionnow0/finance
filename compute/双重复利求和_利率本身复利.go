package main
import (
	"fmt"
	"flag"
)

/*
*    循环法 求x^n
 */
func powerf3(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}

func main() {
	//years := *flag.Int("years", 10, "total years, e.g 10") error
	var years int
	var init_interest_rate , rate_of_increase float64

	flag.IntVar(&years, "years", 10, "total years, e.g 10")
	flag.Float64Var(&init_interest_rate, "rate", 0.04, "first interest rate, e.g 0.02")
	flag.Float64Var(&rate_of_increase, "rate_increase", 0.1, "rate of increase, e.g 0.1")
	flag.Parse()

	fmt.Println("init total money:", 100)
	fmt.Println("years:", years)
	fmt.Println("init_interest_rate:", init_interest_rate)
	fmt.Println("rate_of_increase:", rate_of_increase)

	for i := int(1); i <= years; i++ {
		total := init_interest_rate * (1.0 - powerf3(1.0+rate_of_increase, i)) / (1.0 - (1.0+rate_of_increase))
		fmt.Printf("year %d: {sum: %.2f, thisYearRate: %.1f%%}\n", i, 100 * total, init_interest_rate*powerf3(1.0+rate_of_increase, i-1)*100)
	}
}