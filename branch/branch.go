package main

import (
	"fmt"
	"io/ioutil"
)

// switch
func grade(score int) string {
	g := ""
	switch {
	case score<0 || score>100:
		panic(fmt.Sprintf(
			"Wrong",score))
	case score<60:
		g = "F"
	case score<80:
		g = "C"
	case score<90:
		g = "B"
	case score<100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Printf("%s\n",contents)
	// }

	//条件里赋值 作用域在条件中
	if contents , err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
	fmt.Println(
		grade(90))
	
}