package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString) // &をつけることでそのvariableのlocationを
	// 伝えてますよってことを示せる

	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) { // ここでsにlocationが渡されるよね。
	newValue := "Red"
	*s = newValue // sに*をつけることで今度はsの位置にある値を指すことができる
}