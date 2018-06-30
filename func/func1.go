package main

import (
	"fmt"
)

type user struct {
	id   int
	name string
}

func changeUserName(u user, newName string) user {
	u.name = newName
	return u
}

func changeUserID(u *user, newID int) *user {
	u.id = newID
	return u
}

func printUser(u user) {
	fmt.Println()
	fmt.Printf("User name: [%s]", u.name)
	fmt.Println()
	fmt.Printf("User ID:  [%d]", u.id)
	fmt.Println()
}

type duration int

func main() {

	d := duration(42)
	var dodo duration
	dodo = 88

	var dd int
	dd = 23

	u1 := user{
		id:   23,
		name: "Ahmed",
	}

	// Change the u1 name
	changeUserName(u1, "Hamada")
	printUser(u1)
	fmt.Printf("User Address:  [%p]", &u1)
	fmt.Println()

	// Inspect the change name func return value
	var u2 user
	u2 = changeUserName(u1, "kambas")
	printUser(u2)
	fmt.Printf("User Address:  [%p]", &u2)
	fmt.Println()

	//Change User ID through a reference
	var ux *user
	ux = changeUserID(&u1, 88)
	printUser(u1)
	fmt.Printf("User Address:  [%p]", &u1)
	fmt.Println()
	//Print the returned pointer address
	fmt.Printf("User Address:  [%p]", ux)
	fmt.Println()
}
