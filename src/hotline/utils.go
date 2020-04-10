package hotline

import (
	"fmt"
	"math/big"
	"os"
)

// Factorial - Get factorial big int
func Factorial(x *big.Int) *big.Int {
	result := big.NewInt(1)
	i := big.NewInt(2)

	if !x.IsInt64() {
		fmt.Println("The number is way too big to calculate a factorial")
		os.Exit(84)
	}
    for i.Cmp(x) != 1 {
		result.Mul(result, i)
		i = i.Add(i, big.NewInt(1))
    }
    return result
}

// Pow - big Float
func Pow(a *big.Float, e int64) *big.Float {
    result := big.NewFloat(0.0).Copy(a)
    for i := int64(0); i < e - 1; i++ {
		result = result.Mul(result, a)
    }
    return result
}

// DisplayResult - output format for distributions
func DisplayResult(i int64, res *big.Float) {
	fmt.Printf("%d -> %.3f", i, res)
	if ((i + 1) % 5 == 0 || i == 50) {
		fmt.Printf("\n")
	} else {
		fmt.Printf("\t")
	}
}