// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how functions can return multiple values while using
// named and struct types.
package main

import (
	"encoding/json"
	"fmt"
)

// User is a struct type that declares User information.
type User struct {
	ID   int
	Name string
}

// NewUser Create a new User
func NewUser(name int, value string) *User {
	// retval := new(User)
	retval := User{
		ID:   88,
		Name: "Hamada",
	}
	retval.ID = name
	retval.Name = value
	return &retval
}

func main() {

	// Retrieve the User profile.
	u, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the User profile.
	fmt.Printf("%+v\n", *u)
}

// retrieveUser retrieves the User document for the specified
// User and returns a pointer to a User type value.
func retrieveUser(name string) (*User, error) {

	// Make a call to get the User in a json response.
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of
	// the User struct type.
	var u User
	err = json.Unmarshal([]byte(r), &u)
	return &u, err
}

// GetUser simulates a web call that returns a json
// document for the specified User.
func getUser(name string) (string, error) {
	response := `{"id":1432, "name":"sally"}`
	return response, nil
}
