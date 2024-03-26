package main

import (
	"fmt"
	"reflect"
	"time"
)

func sum(a int, b int) {
	fmt.Println("a,b,c,d,e")
}
func leng(a string) {
	// 定义一个整数切片
	numbers := []int{1, 2, 3, 4, 5}
	// 使用基本的 for循环遍历切片
	for i := 1; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func time_str() {
	// 获取当前时间
	now := time.Now()
	// 格式化时间，使用指定的布局（YYYYMMDDHHMM）
	// 注意：Go的时间格式化布局使用特定的引用时间（Mon Jan 2 15:04:05 MST 2006）来确定各部分的顺序
	formattedTime := now.Format("200601021504")
	// 打印格式化后的时间字符串
	fmt.Println(formattedTime)
}

func main() {
	sum(1, 2)
	fmt.Println(time.DateTime)
	fmt.Println(reflect.TypeOf(time.DateTime))
	fmt.Println(reflect.TypeOf(time.Now()))
	time_str()
}
