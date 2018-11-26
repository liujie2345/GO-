package main

import ("fmt")

// var aa = 3
// var ss = "kkk"
//函数外不能用冒号

var (
		aa = 33
		ss = "kk"
)
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue() {
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,s,b)
}

func variableTypeDeduction() {
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func variableShorter() {
	a,b,c,s := 3, 4, true, "def"
	b = 5
	fmt.Println(a,b,c,s)
}

func main(){
	fmt.Println("hello world")
	variableInitialValue()
	variableZeroValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss)
}