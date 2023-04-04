package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid := saySomething() // これで:をつけて定義することで下で使えるように
	// なっている。これがないとおそらく普通にvar whatWasSaidって定義しないといけないんだろうなあ 

	fmt.Println("The function returned", whatWasSaid)

	anotherSaid, theOtherThingThatWasSaid := doubleSaySomething()
	fmt.Println("The function returned", anotherSaid, theOtherThingThatWasSaid)
}

func saySomething() string {
	return "something"
}

func doubleSaySomething() (string, string) {
	return "something", "else"
}