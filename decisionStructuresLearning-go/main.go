package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println(("Cat is cat"))
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	isNewTrue := false

	if myNum > 99 && !isNewTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isNewTrue {

	} else if myNum == 101 || isNewTrue {

	} else if myNum > 1000 && !isNewTrue {

	}

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println(("cat is something else"))
	}
}