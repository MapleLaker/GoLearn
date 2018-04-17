package main

import "fmt"

//Param is int, return value is int
func max(a int, b int) int{
	if(a>b) {
		return a
	}else{
		return b
	}
	
}

// it has 2 return value
func swap(x, y string) (string, string){
	return y,x
}

func main() {
	var result int
	result = max(50,150)

	fmt.Printf("%d", result)



	var name string
	name = "Monday"

	var author string
	author = "Sunday"
// use this pattern to receive the 2 return value at the same time
	var result1, result2 string = swap(name, author) 

	fmt.Printf("%s\n, %s\n", result1, result2)
}
