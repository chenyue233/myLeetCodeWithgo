package main

import (
	"fmt"
	"time"
)
func MergeSort(data []int)[]int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := MergeSort(data[:middle])
	fmt.Println("left",left)
	right := MergeSort(data[middle:])
	fmt.Println("right",right)
	return left
}
func main() {
	go func() {
		for i:=0;i<10;i++{
			fmt.Println("111",i)
		}
	}()
	go func() {
		for i := 15;i<20;i++ {
			fmt.Println("2222",i)
		}
	}()
	time.Sleep(5*time.Second)
}