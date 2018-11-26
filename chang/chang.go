package main

import ("fmt")
import "math"

func consts() {
	const filename = "abc.txt"
	const a,b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
// 枚举类型

func enmus() {
	const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)

	const (
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}
func main(){
	consts()
	enmus()
}