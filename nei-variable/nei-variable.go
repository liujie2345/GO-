package main

import ("fmt")
import "math"
import "math/cmplx"
/*
	bool string
	(u)int (u)int8 32 64 uintptr
	byte rune(字符)
	float32 64 complex64 128(复数)
*/
//欧拉公式 复数
func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

// go语言没有隐式转化 只有强制显示转化
func triangle() {
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler()
	triangle()
}
