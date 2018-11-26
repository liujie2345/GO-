package main

import (
	"fmt"
)
//数组 切片 容器
//go语言不直接使用数组

// 遍历数组
// for i := 0;i < len(arr3); i++ {
// 	fmt.Println(arr3[i])
// }
func printArray(arr [5]int) {
	arr[0] = 100 //值类型 不会改变
	for i,v := range(arr) {
		fmt.Println(i,v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,2,3} //:=需要初值
	arr3 := [...]int{2,4,6,8,10}
	// 二维数组
	var grid [4][5]int

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	printArray(arr1)
	printArray(arr3)

	fmt.Println(arr1,arr3)  //arr[0] 并没有改变 元素进行拷贝 go语言数组并不是引用传递 *[5]int
}