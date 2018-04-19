package main

import "fmt"

func main() {

	var s1 []int
	printSlice(s1)

	/* 允许追加空切片 */
	s1 = append(s1,0)
	printSlice(s1)	

	/* 同时添加多个元素 */
	s1 = append(s1,1,2,3,4,5)
	printSlice(s1)	

	/* 创建切片 s2 是之前切片的两倍容量*/
	s2 := make([]int, len(s1),cap(s1)*2)
	printSlice(s2)	

	/* 拷贝 s1 的内容到 s2 */
	copy(s2,s1)
	printSlice(s2)

}

func printSlice(number []int) {
	fmt.Printf("len is %d, cap is %d, value is %v\n", len(number), cap(number), number)
}
