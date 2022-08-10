package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ints = make([]int, 0, 5)
	fmt.Println(unsafe.Pointer(&ints))
	i := append(ints, 1, 3, 5, 6)
	fmt.Println(ints)
	fmt.Println(unsafe.Pointer(&i))
	/*fmt.Println(i)
	fmt.Println(len(i))
	fmt.Println(cap(i))*/
}
