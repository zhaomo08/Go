package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为参数
//func main() {
//	ret2 := calc(10, 20, add)
//	fmt.Println(ret2) //30
//}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//func main() {
//	// 将匿名函数保存到变量
//	add := func(x, y int) {
//		fmt.Println(x + y)
//	}
//	add(10, 20) // 通过变量调用匿名函数
//
//	//自执行函数：匿名函数定义完加()直接执行
//	func(x, y int) {
//		fmt.Println(x + y)
//	}(10, 20)
//}

// 闭包:     闭包=函数+引用环境
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

//func main() {
//	var f = adder()
//	fmt.Println(f(10)) //10
//	fmt.Println(f(20)) //30
//	fmt.Println(f(30)) //60
//
//	f1 := adder()
//	fmt.Println(f1(40)) //40
//	fmt.Println(f1(50)) //90
//}

// close	主要用来关闭channel
// len	用来求长度，比如string、array、slice、map、channel
// new	用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
// make	用来分配内存，主要用来分配引用类型，比如chan、map、slice
// append	用来追加元素到数组、slice中
// panic和recover	用来做错误处理

// &（取地址）和*（根据地址取值）

//取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}
