package main

import "fmt"

func main() {
	//指定变量类型，然后赋值
	var i int
	i = 10

	//指定变量类型，同时赋值
	var a int = 10

	//根据值,自行判定变量类型
	var b = 10

	//省略var, 注意 :=左侧的变量不应该是已经声明过的
	c:= 10

	fmt.Printf(" i is %d\n a is %d\n b is %d\n c is %d\n", i, a, b, c)
}
