package hotline

import (
	"fmt"
	"os"
	"math"
	"math/big"
	"time"
	"strconv"
)

type infos struct {
	workingTimeHours int64
	workingTimeSeconds int64
	numberPeople int64
	phoneLine int64
}

func newInfos() infos {
	var infos infos
	infos.workingTimeHours = 8
	infos.workingTimeSeconds = 60 * 60 * infos.workingTimeHours
	infos.numberPeople = 3500
	infos.phoneLine = 25
	return infos
}

func getBinomialCoefficient(n *big.Int, k *big.Int) *big.Int {
	if k.Cmp(n) == 1 {
		fmt.Println("Error: k > n")
		os.Exit(84)
	}

	numerator := Factorial(n)
	subNK := big.NewInt(1).Sub(n, k)
	denominator := big.NewInt(1).Mul(Factorial(k), Factorial(subNK))
	res := big.NewInt(1).Div(numerator, denominator)

	return (res)
}

func computeBinomialDistribution(callTime *big.Int) {
	p := big.NewFloat(1.0)
	infos := newInfos()
	starterTime := float64(time.Now().Nanosecond())
	overload := big.NewFloat(0.0);
	var i int64

	p = p.Quo(big.NewFloat(1.0).SetInt(callTime), big.NewFloat(0.0).SetInt64(infos.workingTimeSeconds))
	
	fmt.Println("Binomial distribution:")
	for i = 0; i <= 50; i++ {
		// Compute
		binCoeff := big.NewFloat(0.0).SetInt(big.NewInt(0).Binomial(infos.numberPeople, i))
		expP := Pow(p, i)
		lastRHS := big.NewFloat(0.0).Sub(big.NewFloat(1), p)
		lastRHSexp := Pow(lastRHS, infos.numberPeople - i)
		res := big.NewFloat(0.0).Mul(binCoeff, expP)
		res = res.Mul(res, lastRHSexp)

		// Display
		DisplayResult(i, res)

		// Overload
		if i <= infos.phoneLine {
			overload = overload.Add(overload, res)
		}
	}
	overload = overload.Sub(big.NewFloat(1.0), overload)
	fmt.Printf("Overload: %.1f", big.NewFloat(0.0).Mul(overload, big.NewFloat(100.0)))
	fmt.Println("%")

	// Computation time
	endTime := float64(time.Now().Nanosecond())
	fmt.Printf("Computation time: %.2f ms\n\n", (endTime - starterTime) / 1000000)    
}

func computePoissonDistribution(callTime string) {
	var p float64
	infos := newInfos()
	starterTime := float64(time.Now().Nanosecond())
	overload := big.NewFloat(0.0);
	var i int64

	callTimeF, err := strconv.ParseFloat(callTime, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(84)
	}
	p = float64(infos.numberPeople) * (callTimeF / float64(infos.workingTimeSeconds))
	
	fmt.Println("Poisson distribution:")
	for i = 0; i <= 50; i++ {
		// Compute
		numerator := big.NewFloat(0.0).Mul(Pow(big.NewFloat(p), i), big.NewFloat(0.0).SetFloat64(math.Exp(-p)))
		denominator := big.NewFloat(0.0).SetInt(Factorial(big.NewInt(i)))
		res := big.NewFloat(0.0).Quo(numerator, denominator)

		// Display
		DisplayResult(i, res)

		// Overload
		if i > infos.phoneLine {
			overload = overload.Add(overload, res)
		}
	}
	fmt.Printf("Overload: %.1f", big.NewFloat(0.0).Mul(overload, big.NewFloat(100.0)))
	fmt.Println("%")

	// Computation time
	endTime := float64(time.Now().Nanosecond())
	fmt.Printf("Computation time: %.2f ms\n", (endTime - starterTime) / 1000000)  
}

// Hotline process
func Hotline() {
	argsWithoutProg := os.Args[1:]
	// Convert args
	first := new(big.Int)
	first, okFirst := first.SetString(argsWithoutProg[0], 10)
	if !okFirst {
		fmt.Println(okFirst)
		os.Exit(84)
	}

	// Redirection
	if len(argsWithoutProg) == 1 {
		computeBinomialDistribution(first)
		computePoissonDistribution(argsWithoutProg[0])
	}
	if len(argsWithoutProg) == 2 {
		second := new(big.Int)
		second, okSecond := second.SetString(argsWithoutProg[1], 10)
		if !okSecond {
			fmt.Println(okSecond)
			os.Exit(84)
		}
		fmt.Println(argsWithoutProg[1] + "-combinations of a set of size " + argsWithoutProg[0] + ":")
		fmt.Println(getBinomialCoefficient(first, second).Text(10))
	}
}