package gotoc

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func Rand() {
	r := C.rand()
	fmt.Println(r)
}
