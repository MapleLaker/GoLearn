package main

import "fmt"

func main() {
	//第一种格式
	for i:=0;i<100;i++{
		fmt.Println(i)
	}

  //第二种格式，类似 while
	var j int = 100;
	for j<200{
		fmt.Println(j)
		j++
	}

  //第三种格式，类似range
	var numbers = [6]int{1,2,3,4,5,6}
	for a,b:= range numbers{
		fmt.Printf("The %d is %d\n",a, b)
	}

}
