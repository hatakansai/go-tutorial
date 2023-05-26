package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.0
	threshold := 1e-10 //閾値の設定
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(prev-z) < threshold { //ひとつ前のzとの差分で絶対値を取り閾値と比較し、変化が小さい場合ループを抜ける
			break
		}
		prev = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
