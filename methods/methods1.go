package main

import "fmt"

type user struct {
	id   int
	name string
}

func (u user) printData() {
	fmt.Println()
	fmt.Printf("User id: %d", u.id)
	fmt.Println()
	fmt.Printf("User name: %s", u.name)
	fmt.Println()
}

func (u user) changeName(name string) {
	u.name = name
}

func (u user) changeID(id int) user {
	u.id = id
	return u
}

func (u *user) changeNameByRef(name string) {
	u.name = name
}

func main() {
	fmt.Printf("Hello World!")
	u1 := user{
		id:   23,
		name: "Ahmed",
	}

	fmt.Println("The original User: ")
	u1.printData()
	fmt.Printf("User address: %p", &u1)
	fmt.Println()

	//Change the user name by val
	u1.changeName("Hamada")
	u1.printData()
	fmt.Printf("User address: %p", &u1)
	fmt.Println()

	//Change name by Ref
	u1.changeNameByRef("kambas")
	u1.printData()
	fmt.Printf("User address: %p", &u1)
	fmt.Println()

	//Change ID by Val
	var u2 user
	u2 = u1.changeID(88)
	u2.printData()
	fmt.Printf("User address: %p", &u2)
	fmt.Println()
	u1.printData()
	fmt.Printf("User address: %p", &u1)
	fmt.Println()
}
