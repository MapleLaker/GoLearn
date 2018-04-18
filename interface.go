package main

import "fmt"

//define interface
type Phone interface{
	Call()
}

//define class which implement the interface 
type Android struct{}
type IPhone struct{}

//class implememtation 
func (anzhuo Android) Call(){
	fmt.Printf("This is Android\n")
}

func (ip IPhone) Call(){
	fmt.Printf("This is IPhone\n")
}

// call 
func main() {

	var phone Phone
	phone = new (Android)
	phone.Call()

	phone = new (IPhone)
	phone.Call()
	
}
