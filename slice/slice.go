package main

import (
	"fmt"
)
//切片
//半开半闭
//slice是对array的view
//可以向后扩展 不能向前扩展

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6}
	s := arr[2:6]

	fmt.Println(s)

	fmt.Println("arr[2:6]",arr[2:6])
	fmt.Println("arr[:6]",arr[:6])
	fmt.Println("arr[2:]",arr[2:])
	fmt.Println("arr[:]",arr[:])

	s1 := arr[:6]
	s2 := arr[:]

	// arter update slice
	updateSlice(s1)
	fmt.Println(s1)  //改变
	fmt.Println(arr)
	updateSlice(s2)
	fmt.Println(s2)  //改变
	fmt.Println(arr)

	// after reslice
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	// extend
	/*
		arr := [...]int{0,1,2,3,4,5,6,7}
		s1 := arr[2:6]
		s2 := s1[3:5]

		s1 = ?
		s2 = ?

		s1 = [2 3 4 5]
		s2 = [5 6]
	*/
}