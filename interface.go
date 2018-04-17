package main

import "fmt"

//define interface
type Phone interface{
	Call()
}

//define Implemtation class 
type Android struct{}

type IPhone struct{}

//class implemation its interface
func (anzhuo Android) Call(){
	fmt.Printf("This is Android\n")
}

func (ip IPhone) Call(){
	fmt.Printf("This is IPhone\n")
}

func main() {

	var phone Phone
	phone = new (Android)
	phone.Call()

	phone = new (IPhone)
	phone.Call()
	
}
