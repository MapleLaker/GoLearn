package main

import "fmt"

func main() {
	
	numbers := []int{1,2,3,4,5}
  //range 数组的数值
	for _,i:= range numbers{
		
		fmt.Println(i)
	}
  //range 数组的序号和大小
	for j, value:= range numbers{
		if value == 3 {
			fmt.Println("Index is ", j)
		}
	}

  //Range, 字符串
	words := "This is a book"
	for _,c := range words{
		fmt.Printf("%c\n", c)
	}

  //Range，Map的key和value
	kvs := map[string]string {"name":"Jack","Author":"ID"}

	for key,value:= range kvs{
		fmt.Printf("The %s is %s\n", key, value)
	}


}
