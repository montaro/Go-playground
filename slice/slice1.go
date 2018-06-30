package main

import "fmt"

func main() {
	var x = 23
	fmt.Println(x)

	// Declare a nil slice of strings.
	var data []string

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Print the data slice addres s
	fmt.Printf("data slice address: [%p]\n", &data)

	// Append ~100k strings to the slice.
	for record := 1; record <= 150; record++ {

		// fmt.Printf("Addr[%p]\n", &data)
		// Use the built-in function append to add to the slice.
		data = append(data, fmt.Sprintf("Rec: %d", record))

		if data != nil {
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d]\n",
				&data[0],
				record,
				cap(data))
		}

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {

			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}

	fmt.Printf("Addr[%p]\n", &data)
}
