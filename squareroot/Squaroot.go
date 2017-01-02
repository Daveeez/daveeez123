package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, int) {
	z := 0.9
	zprev := z - .1
	zdiff := 1.0
	j := 1
	for i := 1; zdiff > 0.000000000000001; i++ {
		zprev = z
		z = z - (z*z-x)/(2*z)
		zdiff = z - zprev
		if zdiff < 0 {
			zdiff = zprev - z
		}
		j = i
	}
	return z, j
}

func main() {
	fmt.Println(Sqrt(2))
}
