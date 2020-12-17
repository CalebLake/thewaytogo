package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// hello
func main() {
	// 所有声明的变量必须使用，如果没有使用必须删去声明
	quantity := 4
	length, width := 1.2, 2.4

	var a = "hello"
	fmt.Println(quantity)
	fmt.Println(length*width, "square meters")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
	fmt.Println(a)

	// 循环
	var now time.Time = time.Now()
	for x := 1; x < 3; x++ {
		fmt.Println(x)
	}
	fmt.Println(now)

	//指针
	amount := 6
	double(&amount)
	fmt.Println(amount)

	//切片 数组
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Println(string(b[1:4]), b[1:4])
	fmt.Println(string(b[:2]), b[:2])
	fmt.Println(len(b))

}

func double(number *int) {
	*number *= 2
}
