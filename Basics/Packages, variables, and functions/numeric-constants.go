/*
package main

import "fmt"

const (

	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big int64 = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99

)

func needInt(x int64) int64 { return x*10 + 1 }

	func needFloat(x float64) float64 {
		return x * 0.1
	}

	func main() {
		fmt.Println(needInt(int64(Small)))
		fmt.Println(needInt(int64(Big)))
		fmt.Println(needFloat(float64(Small)))
		fmt.Println(needFloat(float64(Big)))
	}
*/
//上記ではオーバーフローの対応ができないためbigを用いる
package main

import (
	"fmt"
	"math/big"
)

var (
	// Create a huge number by raising 2 to the power of 100.
	Big   = new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil)
	Small = new(big.Int).Rsh(Big, 99)
)

func needInt(x *big.Int) *big.Int {
	return new(big.Int).Add(new(big.Int).Mul(x, big.NewInt(10)), big.NewInt(1))
}

func needFloat(x *big.Int) *big.Float {
	return new(big.Float).Mul(new(big.Float).SetInt(x), big.NewFloat(0.1))
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
