package main

import (
	"runtime"
	"reflect"
	"fmt"
)

func eval(a,b int, op string) (int,error) {
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		// return a/b
		q,_ := div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("wrong: %s" , op)
	}
}

// 返回除值 求余
func div(a,b int) (q,r int) {
	// q = a/b
	// r = a%b
	return a / b, a % b
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d.%d)",opName,a,b)
	return op(a,b)
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(13,3,"x"))
	q, r := div(13,3)
	fmt.Println(q,r)
	fmt.Println(sum(1,2,3,4,5))
}