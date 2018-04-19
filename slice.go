package main

import "fmt"

func main() {
	
	var numbers = make([]int, 3, 5)

	printSlice(numbers)

}

//打印数组，%v
func printSlice(number []int) {
	fmt.Printf("len is %d, cap is %d, value is %v", len(number), cap(number), number)
}
