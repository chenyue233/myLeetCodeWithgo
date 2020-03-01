package main

import (
	"fmt"
)

func main() {
	a := [][]int{{1, 2, 3}, {123, 12, 1}}
	a[0] = append(a[0], 1)
	fmt.Println(a)
}
