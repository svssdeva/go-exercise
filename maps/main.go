package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//maps
	// Maps in Go are key-value pairs, similar to dictionaries in Python or hash tables in other languages.
	// They are unordered collections and can be used to store data in a way that allows for fast lookups, additions, and deletions.
	// A map is defined using the `map` keyword followed by the key type and value type in square brackets.
	// Example: Creating a map of string keys and int values
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	// Accessing values
	fmt.Println(m["apple"]) // Output: 5
	// Adding a new key-value pair
	m["grape"] = 7
	// Deleting a key-value pair
	delete(m, "banana")
	// Iterating over the map
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	//map for json and api
	jsonMap := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	// Adding a nested map
	jsonMap["address"] = map[string]string{
		"street": "123 Main St",
		"city":   "New York",
	}
	// Accessing nested map values
	fmt.Println(jsonMap["address"].(map[string]string)["city"]) // Output: New York
	// Converting map to JSON string (using encoding/json package)
	// import "encoding/json"
	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData)) // Output: {"name":"John","age":30,"address":{"street":"123 Main St","city":"New York"}}
	// Unmarshalling JSON string back to map
	var unmarshalledMap map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshalledMap)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(unmarshalledMap) // Output: map[name:John age:30 address:map[street:123 Main St city:New York]]

	//make vs map
	// The `make` function is used to create slices, maps, and channels in Go.
	madeMap := make(map[string]int) // Creating an empty map using make
	madeMap["key1"] = 1
	madeMap["key2"] = 2
	fmt.Println(madeMap) // Output: map[key1:1 key2:2]
	// The `map` keyword is used to define a map literal, which is a predefined set of key-value pairs.
	// Example of map literal
	literalMap := map[string]int{
		"keyA": 10,
		"keyB": 20,
	}
	fmt.Println(literalMap) // Output: map[keyA:10 keyB:20]

	//make vs new
	// The `make` function is used to create slices, maps, and channels, while the `new` function is used to allocate memory for a variable of a specific type.
	newMap := new(map[string]int)  // Allocating memory for a map using new
	*newMap = make(map[string]int) // Initializing the map using make
	(*newMap)["keyX"] = 100
	(*newMap)["keyY"] = 200
	fmt.Println(*newMap) // Output: map[keyX:100 keyY:200]
	// The `new` function returns a pointer to the allocated memory, while `make` returns the actual map value.

	//type definition (not alias) to allow methods
	type StringIntMap map[string]int // Creating a named type for a map
	aliasMap := StringIntMap{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(aliasMap) // Output: {"key1":1,"key2":2}
}
