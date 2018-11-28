package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
	golang中的数据类型
	基本数据类型
		数值型：
			int（2的多少次幂与系统有关32/64）、int8(byte)、int16、int32(rune)、int64、
			范围：intx -> （-2的x幂 ~ 2的x幂-1）
			（无符号类型）uint、uint8、uint16、uint32、uint64、byte
			范围：uintx -> （0~ 2的x幂-1）
		小数：
			单精度：float32
			双精度：float64
		字符型：byte（0~255）
		布尔型：bool
		字符串类型：string
	派生数据类型
		指针：Pointer
		数组
		结构体：struct
		管道：channel
		函数
		切片：slice
		接口：interface
		map
*/
func main() {
	// foo()
	// bar()
	// likeChar()
	// boolDemo()
	// stringDemo()
	// zeroValue()
	// transForm()
	// dataTransExercise()
	//otherDataTypesToString()
	stringToOthers()
}

// 整型
func foo() {

	var num1 int8 = -128
	var num2 int8 = 127 // constant 128 overflows int8b
	fmt.Println(num1, num2)
	var num3 uint8 = 0
	var num4 uint8 = 255
	fmt.Println(num3, num4)
	// int rune（相当于int32） byte(相当于无符号的int8)
	var num5 int = 0
	var num6 uint = 1
	var num7 byte = 'a' // 97
	var num8 rune = 255
	fmt.Println(num5, num6, num7, num8)
	// 整型使用细节
	// 默认int
	num9 := 100
	var num10 byte = 'z'
	fmt.Printf("num9值 :%d, 类型：%T\n", num9, num9)
	fmt.Printf("num10值 :%d, 类型：%T\n", num10, num10)
	// 查看某个变量占用的字节大小
	fmt.Printf("num10值 :%d, 类型：%T, 占用的字节数为：%d\n", num10, num10, unsafe.Sizeof(num10))
	// 保小不保大，不要浪费资源
	var age byte = 10 // 0 ~ 255
	var count int64
	fmt.Println(age, count)
}

// 浮点型：推荐使用float64
func bar() {
	// 1.0123457
	var price1 float32 = 1.0123456789
	// 1.0123456789
	var price2 float64 = 1.0123456789
	// 浮点数无论单精度还是双精度都是由符号位、指数位、尾号位组成的，可能造成精度损失
	fmt.Println(price1, price2)
	// 默认为 float64
	var price3 = 1.01
	fmt.Printf("price3的数据类型为 %T，占用的字节数为：%d \n", price3, unsafe.Sizeof(price3))
	// 十进制和科学计数法形式
	num1 := 1.234e2
	num2 := 1.234E2
	num3 := 1.234E-2
	num4 := .123
	fmt.Println(num1, num2, num3, num4) // 123.4 123.4 0.01234 0.123
}

// go中没有字符的单独表示，字符和字节一样，使用byte表示，字符串由byte组成
// go中字符型本质上是整型
func likeChar() {
	// 直接输出byte的值，实际上是对应ascii码的值
	var ch1 byte = 'a'
	var ch2 byte = '0'
	//97 0
	fmt.Println(ch1, ch2)
	// 需要格式化输出
	fmt.Printf("ch1=%c,ch2=%c,类型为：%T", ch1, ch2, ch1)
	//var ch3 byte = '朱'
	//fmt.Println(ch3)
	// 如果char的码值小于255，用byte类型，过大用int、rune类型
	var ch4 rune = '朱'
	fmt.Println(ch4)
	fmt.Printf("字符表示为：%c， ascii码值为：%d\n", ch4, ch4) // 字符表示为：朱， ascii码值为：26417 不在0~255之间。不可使用byte(uint8)
	// 注意：unicode编码的一种实现为utf-8， unicode为ascii码的拓展
	var ch5 byte = '\n'
	fmt.Printf("字符表示为：%c， ascii码值为：%d\n", ch5, ch5)
	// go中字符型可以进行运算
	var num = 1 + 'a'
	fmt.Println("num = ", num)
}

// bool 类型只能true或false， 占用一个字节
func boolDemo() {
	b := false
	fmt.Printf("值为：%v，占用字节数：%d", b, unsafe.Sizeof(b))
}

// 字符串类型
func stringDemo() {
	str1 := "哈哈 嘻嘻 呵呵"
	fmt.Println(str1, str1[0])
	str2 := "aa\n\""
	fmt.Println(str2)
	// 使用``定义字符串
	str3 := `aa\n""`
	fmt.Println(str3)
	// 字符串的拼接
	str4 := str1 + str3
	str4 += str4
	fmt.Println(str4)
}

// 基本数据类型的默认值（零值）
func zeroValue() {
	var a int
	var b byte
	var c rune
	var d uint
	var e float32
	var f float64
	var g bool
	var h string
	fmt.Println(a, b, c, d, e, f, g, h) // 0 0 0 0 0 0 false ""
}

// go 中基本数据类型转换特点：没有自动类型转换，必须显示转换（强转）
func transForm() {
	// int64 转 float64 -->转高精度
	var num1 int64 = 100
	var num2 float64 = float64(num1)
	fmt.Printf("num2:%v, type:%T\n", num2, num2)
	// int64 转 int8  -->转低精度(注意：结果可能不正确)
	var num3 int8 = int8(num1)
	fmt.Printf("num3:%v, type:%T\n", num3, num3)
	// float64 转 int64
	var num4 float64 = 3.3333
	var num5 int64 = int64(num4)
	fmt.Printf("num5:%v, type:%T\n", num5, num5)
}
func dataTransExercise() {
	var n1 int32
	var n2 int64
	var n3 int8
	n4 := n1 + 20
	// ‘20’ 会自动类型推断成n1的数据类型
	fmt.Printf("n4:%v,type:%T\n", n4, n4) // n4:20,type:int32
	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println(n2, n3)
	n1 = 1
	n3 = int8(n1) + 127 // 运行报错
	// n3 = int8(n1) + 128 // 编译报错
	// 向低精度类型转换的时候一定要注意数据的范围有没有超出类型的范围
	fmt.Println(n3)
}

// 其他基本数据类型转换为string 类型
func otherDataTypesToString() {
	var a int = 20
	var b float64 = 3.3333
	var c bool = true
	var d byte = 'a'
	var e float32 = 2.21
	// （方式一）
	aResp := fmt.Sprintf("%d", a)
	bResp := fmt.Sprintf("%f", b)
	cResp := fmt.Sprintf("%t", c)
	dResp := fmt.Sprintf("%c", d)
	eResp := fmt.Sprintf("%f", e)
	fmt.Printf("%q,%q,%q,%q,%q\n", aResp, bResp, cResp, dResp, eResp) // "20","3.333300","true","a","2.210000"
	// (方式二)
	aa := strconv.FormatInt(int64(a), 10)
	bb := strconv.FormatFloat(b, 'f', 10, 64)
	cc := strconv.FormatBool(c)
	dd := strconv.FormatInt(int64(d), 10)
	ee := strconv.FormatFloat(float64(e), 'f', 10, 64)
	fmt.Printf("%q,%q,%q,%q,%q\n", aa, bb, cc, dd, ee) // "20","3.3333000000","true","97","2.2100000381"
}

// string 类型转换为其他基本数据类型
func stringToOthers() {
	// 若无法转换，不会报错，会转为默认值
	var a string = "123"
	result, err := strconv.ParseUint(a, 10, 64)
	if err == nil {
		fmt.Printf("%d\n", result)
	}

	var b string = "a"
	result2, _ := strconv.ParseUint(b, 10, 64)
	fmt.Printf("%d", result2)
}
