package main

import "fmt"

func main() {
	//创建slice
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	//s1 := [3]int: It is wrong as init != declaration 
	//仅仅声明，不赋值
	var s1 []int
	printSlice(s1)
	if s1 == nil {
		fmt.Println("This is empty container\n")
	}

	//创建 切片
	s2 := []int{1,2,3,4,5,6,7,8,9}
	/* 打印原始切片 */ 
	fmt.Printf("s2 is %v\n", s2)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Printf("s2[1:4] is %v\n", s2[1:4])

	/* 默认下限为 0*/ 
	fmt.Printf("s2[:4] is %v\n", s2[:4])

	/* 默认上限为 len(s)*/ 
	fmt.Printf("s2[1:] is %v\n", s2[1:])

	/* 创建 子切片 */
	s3 := s2[2:5]
	printSlice(s3)


}

func printSlice(number []int) {
	fmt.Printf("len is %d, cap is %d, value is %v\n", len(number), cap(number), number)
}
