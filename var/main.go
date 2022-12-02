package main

import "fmt"

var a string
var b int
var c bool
/*var{
	a string
	b int
	c bool
}
*/
func main() {
	a = "hello,world!"
	b = 10
	c = true
	/*fmt.Println(a,b,c)*/
	fmt.Printf("a:%s\n",a)
	fmt.Println("b:",b)
	fmt.Print("c:",c)
}