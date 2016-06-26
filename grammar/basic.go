package grammar

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

type ByteSize float64

var x string = "Hello world"

func Hello() {
	fmt.Println("Hello World.")
}

func Typesize() {
	i := 5
	var j int
	fmt.Println("i is : ", i)
	fmt.Println("j is : ", j)
	fmt.Printf("Size i : %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("Size j : %d bytes\n", unsafe.Sizeof(j))
}

func Enum() {
	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Printf("%32.f\n", KB)
	fmt.Printf("%32.f\n", MB)
	fmt.Printf("%32.f\n", GB)
	fmt.Printf("%32.f\n", TB)
	fmt.Printf("%32.f\n", PB)
	fmt.Printf("%32.f\n", EB)
	fmt.Printf("%32.f\n", ZB)
	fmt.Printf("%32.f\n", YB)
}

func Scope() {
	fmt.Println(x)
}

func Scanf() {
	fmt.Print("Enter a number :")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func _zero(xPtr *int) {
	fmt.Printf("xPtr : %p\n", xPtr)
	*xPtr = 0
}

func _one(xPtr *int) {
	fmt.Printf("xPtr : %p\n", xPtr)
	*xPtr = 1
}

func Pointer() {
	x := 5
	fmt.Printf("address of x : %p\n", &x)
	_zero(&x)
	fmt.Println("x : ", x)

	xPtr := new(int)
	fmt.Printf("address of x : %p\n", xPtr)
	_one(xPtr)
	fmt.Println("value of xPtr : ", *xPtr)
}

func Average() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println("Average is : ", _func_average(xs))
}

func _func_average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return (total / float64(len(xs)))
}

func Multi_return() {
	fmt.Println(_func_multi_return(13, 6))
	fmt.Println(_func_multi_return(5, 0))
}

func _func_multi_return(x int, y int) (ret int, err error) {
	if y == 0 {
		ret = int(math.NaN())
		err = errors.New("haha")
	} else {
		ret = x / y
	}
	return
}

func First_class() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "Hello World! My name is Hello!"))
}

func Anonymous_func() {
	x := 5
	fn := func() {
		fmt.Println("x is ", x)
	}
	fn()
	x++
	fn()
}
