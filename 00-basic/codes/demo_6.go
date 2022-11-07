//demo_6.go

package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArrByPointer(&arr)
	fmt.Println(arr)
}

func modifyArrByPointer(a *[5]int) {
	a[1] = 20
}
