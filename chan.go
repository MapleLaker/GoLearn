package main

import (
"fmt"
)

func sum(a []int, c chan int) {
     
	var total = 0
  for _, i := range a{
    	total += i
  }

  fmt.Println(total)

  c <- total
}

func main() {
	
	a := []int{1,2,3,4,5,6,7}
	c := make(chan int)

	go sum(a,c)

	//time.Sleep(2 * 1e9)

  //同步，直到chan c 给 x 赋值后，才会继续执行
	x := <-c

	fmt.Println(x)
	
	
	//buffer chan
	//e := make(chan int, 1) - dead lock : c <- 2 is alwags waiting as y := <- e is not executed
	e := make(chan int, 2)

	e <- 1
	e <- 2

	y := <- e
	z := <- e
	fmt.Println(y, z)
	
}
