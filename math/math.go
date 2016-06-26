package math

import (
	"fmt"
)

func Sqrt() {
	fmt.Println("Hello World. Sqrt(2)=", sqrt(2))
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
