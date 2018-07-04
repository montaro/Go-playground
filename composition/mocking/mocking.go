// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how you can personally mock concrete types when
// you need to for your own packages or tests.
package main

import (
	"fmt"
)

// publisher is an interface to allow this package to mock the pubsub
// package support.
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

type pusher interface {
	Publish(key string, v interface{}) error
}

// service that uses an Event Bus to push events
func serviceX(puppy publisher, event string) {
	fmt.Println()
	fmt.Println("Service X started to execute...")
	fmt.Printf("Service X will push the event: %s", event)
	fmt.Println("")
	puppy.Publish("serviceX", event)
	fmt.Println("")
	fmt.Println("Service X shutdown...")
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	fmt.Println()
	fmt.Printf("%T I will publish to the key: %s with value: %s\n", m, key, v)
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {

	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	fmt.Println("I got: " + key)
	return nil
}

// mock is a concrete type to help support the mocking of the pubsub package.
type pushmock struct{}

// Publish implements the publisher interface for the mock.
func (m *pushmock) Publish(key string, v interface{}) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	fmt.Printf("%T I will push the key: %s with value: %s\n", m, key, v)
	return nil
}

func main() {

	// Create a slice of publisher interface values. Assign
	// the address of a pubsub.PubSub value and the address of
	// a mock value.
	// pubs := []publisher{
	// 	pubsub.New("localhost"),
	// 	&mock{},
	// }
	// pushers := []pusher{
	// 	pubsub.New("localhost"),
	// 	&pushmock{},
	// }
	// Range over the interface value to see how the publisher
	// interface provides the level of decoupling the user needs.
	// The pubsub package did not need to provide the interface type.
	// for _, p := range pubs {
	// 	p.Publish("key", "value")
	// }

	// for _, p := range pushers {
	// 	p.Publish("key", "value")
	// }

	// Use the service X to publish an event with a concrete Event Bus
	// var realPubsub = pubsub.New("example.com")
	// serviceX(realPubsub, "Real Event")
	var mockPubsub = mock{}
	serviceX(&mockPubsub, "Test Event")
}
