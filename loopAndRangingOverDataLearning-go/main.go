package main

import "log"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// for i, animal := range animals { // これで勝手にiにindexが入っているみたい
	// 	log.Println(i, animal)
	// }

	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// newAnimals := make(map[string]string)

	// newAnimals["dog"] = "Fido"
	// newAnimals["cat"] = "Fluffy"

	// for animalType, animal := range newAnimals { // 今回はmapなので
	// 	// animalTypeには単にkeyが入っている
	// 	log.Println(animalType, animal)
	// }

	// var firstLine = "Once upon a midnight dreary"

	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User

	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}