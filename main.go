package main

import (
	"fmt"
)

func main() {
	//fmt.Println(atoi.MyAtoi(str))
	str := "7sss"
	//a := "8"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(str[i])
	// }
	res := 1
	a := int(str[0])
	s := res*10 + int(str[0])
	fmt.Println(a, s)
}
