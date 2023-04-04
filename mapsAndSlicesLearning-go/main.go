package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
	// myMap := make(map[string]string)

	// myMap["dog"] = "Samson"

	// myMap["other-dog"] = "Cassie"

	// myMap["dog"] = "fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// どうやらmapはsortできなくてslicesがsortできるみたい。
	// ちなみに通常のprogrammingでいうarrayがslicesにあたるんだそうです

	// myOtherMap := make(map[string]int)

	// myOtherMap["First"] = 1
	// myOtherMap["Second"] = 2

	// log.Println(myOtherMap["First"])
	// log.Println(myOtherMap["Second"])

	// myAnothermap := make(map[string]User)
	// me := User {
	// 	FirstName: "Trevor",
	// 	LastName: "Sawler",
	// }

	// myAnothermap["me"] = me
	// log.Println(myAnothermap["me"].FirstName)

	// var mySlice []string
	// mySlice = append(mySlice, "Tatsu")
	// mySlice = append(mySlice, "John")
	// log.Println(mySlice)

	var myIntSlice []int
	myIntSlice = append(myIntSlice, 2)
	myIntSlice = append(myIntSlice, 1)
	myIntSlice = append(myIntSlice, 3)

	log.Println(myIntSlice)

	sort.Ints(myIntSlice)

	log.Println(myIntSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:2]) // 0以上2未満の範囲ってこと
}