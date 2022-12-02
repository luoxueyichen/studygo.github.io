package main

import "fmt"

//var a string
//var b int
//var c bool
/*批量声明变量
var{
	a string
	b int
	c bool
}
*/
func main() {
	/*
	a = "hello,world!"
	b = 10
	c = true
	fmt.Println(a,b,c)
	fmt.Printf("a:%s\n",a)
	fmt.Println("b:",b)
	fmt.Print("c:",c)
	*/
	/*
	直接声明变量以及赋值
	var s1 string = "woshiyigezifuchuan"
	var s2 int = 20
	var s3 bool = false
	fmt.Println(s1,s2,s3)
	*/
	//简短声明变量（:=），只能在函数内部使用
	s1:= "woshiyigezifuchuan"
	s2:= 20
	s3:= false
	fmt.Println(s1,s2,s3)
}