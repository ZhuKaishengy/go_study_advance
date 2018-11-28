package main

import (
	"fmt"
)

func main() {

	//fmt.Println	("hello world")
	// // 转义字符
	//fmt.Println("haha \t haha")
	//fmt.Println("D:\\Workspaces\\GolangWorkspace\\src\\go_code\\project01\\main")
	//hi();
	//testFormat()
	//time.Sleep(5 * time.Second)
	//myCalc()
	//testAttr()
	// testVariable()
	plusOperator()
}
func hi() {
	fmt.Println("hi world")
	var name = "zks"
	fmt.Println(name)
}

/*
	练习转义字符
*/
func testFormat() {
	fmt.Println("name\tage\tprovince\tcity\nSam\t21\tLN\tSY")
	fmt.Println("姓名\t年龄\t省\t市\nSam\t21\tLN\tSY")
}

// go 中函数的返回值可以是多个
func calc(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func myCalc() {
	// 说明int 类型初始值为0
	var num1 int
	var num2 int
	// 0 0
	fmt.Println(num1, num2)
	num1 = 20
	num2 = 10
	sum, sub := calc(num1, num2)
	fmt.Println(sum, sub)
}

/*
	golang使用变量的三种方式
	1、指定变量类型，声明但不赋值，使用默认值
	2、不指定变量类型，go中有自动类型推断
	3、省略var，使用":="，这是声明+赋值
*/

func testAttr() {
	// 1、指定变量类型，声明但不赋值，使用默认值
	var name string
	fmt.Println(name)
	// 2、不指定变量类型，go中有自动类型推断，能推断出来尽量省略
	var age = 10
	fmt.Println(age)
	// 3、省略var，使用":="，这是声明+赋值
	desc := "a good boy"
	fmt.Println(desc)
}

/*
	golang中的全局变量和局部变量
	在方法外声明的为全局变量，
	在方法内声明的为局部变量
*/

// 全局变量不能使用":=" 的方式初始化
var name1, name2 = "sjx", "haha"
var (
	age  = 12
	desc = "a good girl"
)

func testVariable() {
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(age)
	fmt.Println(desc)
	// 多个局部变量的声明方式
	person, sex := "person1", true
	fmt.Println(person, sex)
}

// golang "+" 使用
func plusOperator() {
	var a int = 1
	var b int = 2
	fmt.Println(a + b)
	str1 := "a"
	str2 := "b"
	fmt.Println(str1 + str2)
}
