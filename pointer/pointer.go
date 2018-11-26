package main

import (
	"fmt"
)

//交换值
func swap(a,b *int) {
	*b,*a = *a,*b
}

//第二种方法
// func swap(a,b int) (int, int) {
// 	return b,a
// }
func main() {
	a, b := 3, 4
	swap(&a,&b)
	fmt.Println(a,b)
}