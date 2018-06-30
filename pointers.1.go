package main

import "fmt"

func main() {
	// i, j := 42, 2701

	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i
	// fmt.Println(*p) // read i through the pointer
	// fmt.Println(p) // read the pointer p value (i value address)
	// fmt.Println(&p) // read the pointer p's itself address
	// fmt.Println(&i) // read the pointer i's itself address
	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j

	type user struct {
		ID   int
		Name string
	}

	var u user
	var uu *user
	u.ID = 23
	u.Name = "Ahmed"
	var hamada user
	hamada.ID = 88
	hamada.Name = "Hamada"
	uu = &u

	fmt.Println(u)
	// fmt.Println(*u)
	fmt.Println(&uu)
	fmt.Println(*uu)
}
