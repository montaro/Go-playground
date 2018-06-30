// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the behavior of the for range and
// how memory for an array is contiguous.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Declare an array of 5 strings initialized with values.
	five := [5]string{"日本語; charset=utf-8", "#", "世界", "Doug", "Hello, \x90\xA2\x8A\x45"}
	// five := [5]byte{1, 2, 3, 4, 5}

	// Iterate over the array displaying the value and
	// address of each element.
	for i, v := range five {
		fmt.Printf("Value[%v]\tAddress[%p] IndexAddr[%p]\tLength of string[%v]\n", v, &v, &five[i], len(v))
	}

	x := "Ahmed ElRefaey Hamouda"
	fmt.Println("\n")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(&x)
	fmt.Println(len(*&x))
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(five))
	fmt.Println("\n")

	s1 := "foo"
	s2 := "foobar"

	fmt.Printf("s1 size: %T, %d\n", s1, unsafe.Sizeof(s1))
	fmt.Printf("s2 size: %T, %d\n", s2, unsafe.Sizeof(s2))
	fmt.Printf("s1 len: %T, %d\n", s1, len(s1))
	fmt.Printf("s2 len: %T, %d\n", s2, len(s2))
	fmt.Printf("s1 address: %v, %v\n", s1, &s1)
	fmt.Printf("s2 address: %v, %v\n", s2, &s2)

	fmt.Println("\n")
	s := "Hellõ World"

	fmt.Println("len(s)", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%T ", s[i])
	}

	r := []rune(s) // type conversion

	fmt.Println("len(r)", len(r))

	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}

	fmt.Println("")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%v ", r[i])
	}

	fmt.Println("")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%x ", r[i])
	}

	fmt.Println("")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%T ", r[i])
	}
}
