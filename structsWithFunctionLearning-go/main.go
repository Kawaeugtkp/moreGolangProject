package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string { // これはreceiverといって、
	// この表記にすることで、この関数をmyStructの要素として使うことができるから、
	// myStructのインスタンスはこのprintFirstNameを呼べる
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())

}