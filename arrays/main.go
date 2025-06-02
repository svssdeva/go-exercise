package main

import "fmt"

func main() {
	hobbits := [5]string{"Frodo", "Sam", "Merry", "Pippin", "Bilbo"}
	fmt.Println("Hobbits:", hobbits)

	// Accessing elements
	fmt.Println("First hobbit:", hobbits[0])
	fmt.Println("Last hobbit:", hobbits[len(hobbits)-1])

	// Slicing
	fmt.Println("Middle hobbits:", hobbits[1:3])
	fmt.Println("First 3 hobbits:", hobbits[:3])
	fmt.Println("Last 2 hobbits:", hobbits[len(hobbits)-2:])

	// Length of the array
	fmt.Println("Number of hobbits:", len(hobbits))

	// Iterating over the array
	for i, hobbit := range hobbits {
		fmt.Printf("Hobbit %d: %s\n", i+1, hobbit)
	}

	// Modifying an element
	hobbits[0] = "Bilbo"
	fmt.Println("Modified hobbits:", hobbits)

	// Creating a new array
	newHobbits := [3]string{"Gandalf", "Aragorn", "Legolas"}
	fmt.Println("New hobbits:", newHobbits)

	// Combining arrays (not directly possible, but can be done with slices)
	combinedHobbits := append(hobbits[:2], newHobbits[:]...)
	fmt.Println("Combined hobbits:", combinedHobbits)
	fmt.Println("Combined hobbits length:", len(combinedHobbits))
	fmt.Println("Combined hobbits capacity:", cap(combinedHobbits))
	fmt.Println("Combined hobbits as array:", [5]string{combinedHobbits[0], combinedHobbits[1], combinedHobbits[2], combinedHobbits[3], combinedHobbits[4]})
	fmt.Println("Combined hobbits as slice:", combinedHobbits[:5])
	fmt.Println("Combined hobbits as slice length:", len(combinedHobbits[:5]))

	// Demonstrating zero value of an array
	zeroHobbits := [5]string{}
	fmt.Println("Zero value hobbits:", zeroHobbits)
	fmt.Println("Zero value hobbits length:", len(zeroHobbits))
	fmt.Println("Zero value hobbits capacity:", cap(zeroHobbits))

	// Demonstrating array comparison (not allowed in Go)
	// fmt.Println("Comparing arrays:", hobbits == newHobbits) // This will cause a compile error

	// Demonstrating array equality (not allowed in Go)
	// fmt.Println("Are hobbits equal to newHobbits?", hobbits == newHobbits) // This will cause a compile error

	// Demonstrating array type conversion (not allowed in Go)
	// fmt.Println("Hobbits as interface:", interface{}(hobbits)) // This will cause a compile error
	fmt.Println("Hobbits as interface:", interface{}(hobbits[:])) // This works because slices can be converted to interface{}
	fmt.Println("Hobbits as interface length:", len(interface{}(hobbits[:]).([]string)))
	fmt.Println("Hobbits as interface capacity:", cap(interface{}(hobbits[:]).([]string)))
	fmt.Println("Hobbits as interface type:", fmt.Sprintf("%T", interface{}(hobbits[:])))
	fmt.Println("Hobbits as interface value:", interface{}(hobbits[:]))
	fmt.Println("Hobbits as interface value type:", fmt.Sprintf("%T", interface{}(hobbits[:])))
	fmt.Println("Hobbits as interface value length:", len(interface{}(hobbits[:]).([]string)))
	fmt.Println("Hobbits as interface value capacity:", cap(interface{}(hobbits[:]).([]string)))
	fmt.Println("Hobbits as interface value elements:", interface{}(hobbits[:]).([]string)[0:5])

	// Demonstrating array of arrays
	nestedHobbits := [2][3]string{
		{"Frodo", "Sam", "Merry"},
		{"Pippin", "Bilbo", "Gandalf"},
	}
	fmt.Println("Nested hobbits:", nestedHobbits)
	fmt.Println("First nested hobbit:", nestedHobbits[0][0])
	fmt.Println("Second nested hobbit:", nestedHobbits[1][1])
	fmt.Println("Length of nested hobbits:", len(nestedHobbits))
	fmt.Println("Length of first nested hobbits array:", len(nestedHobbits[0]))
	fmt.Println("Length of second nested hobbits array:", len(nestedHobbits[1]))
	fmt.Println("Nested hobbits as interface:", interface{}(nestedHobbits))
	fmt.Println("Nested hobbits as interface type:", fmt.Sprintf("%T", interface{}(nestedHobbits)))
	fmt.Println("Nested hobbits as interface value:", interface{}(nestedHobbits))
	fmt.Println("Nested hobbits as interface value type:", fmt.Sprintf("%T", interface{}(nestedHobbits)))

	// Demonstrating multidimensional arrays
	fmt.Println("Multidimensional array of hobbits:")
	multiHobbits := [2][2][2]string{
		{
			{"Frodo", "Sam"},
			{"Merry", "Pippin"},
		},
		{
			{"Bilbo", "Gandalf"},
			{"Aragorn", "Legolas"},
		},
	}
	fmt.Println(multiHobbits)
	for i, row := range multiHobbits {
		for j, col := range row {
			for k, hobbit := range col {
				fmt.Printf("Hobbit at [%d][%d][%d]: %s\n", i, j, k, hobbit)
			}
		}
	}
	fmt.Println("Length of multidimensional hobbits:", len(multiHobbits))

	// Demonstrating array of slices
	sliceOfHobbits := [][]string{
		{"Frodo", "Sam"},
		{"Merry", "Pippin"},
		{"Bilbo", "Gandalf"},
		{"Aragorn", "Legolas"},
	}
	fmt.Println("Slice of hobbits:", sliceOfHobbits)
	for i, slice := range sliceOfHobbits {
		fmt.Printf("Slice %d: %v\n", i, slice)
	}
	fmt.Println("Length of slice of hobbits:", len(sliceOfHobbits))

	// Demonstrating array of interfaces
	interfaceHobbits := []interface{}{
		"Frodo",
		42,
		3.14,
		true,
		[]string{"Sam", "Merry"},
		map[string]string{"Pippin": "Bilbo"},
	}
	fmt.Println("Interface hobbits:", interfaceHobbits)
	fmt.Println("Interface hobbits as interface:", interface{}(interfaceHobbits))
	fmt.Println("Interface hobbits as interface type:", fmt.Sprintf("%T", interface{}(interfaceHobbits)))
	fmt.Println("Interface hobbits as interface value:", interface{}(interfaceHobbits))
	fmt.Println("Interface hobbits as interface value type:", fmt.Sprintf("%T", interface{}(interfaceHobbits)))
	fmt.Println("Interface hobbits as interface value length:", len(interface{}(interfaceHobbits).([]interface{})))
	fmt.Println("Interface hobbits as interface value capacity:", cap(interface{}(interfaceHobbits).([]interface{})))
	fmt.Println("Interface hobbits as interface value elements:", interface{}(interfaceHobbits).([]interface{})[0:5])

	fmt.Println("End of array demonstration.")
}
