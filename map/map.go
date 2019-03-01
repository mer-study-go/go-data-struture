package main

import "fmt"

/*
	+--------------------------+
	| key   | value            |
	|-------|------------------|
	| name  | Mer              |
	| email | mer@mercari.com  |
	| role  | developer        |
	+-------|------------------+

*/
func main() {
	// map[key]value
	user := map[string]string{
		"name":  "Mer",
		"email": "mer@mercari.com",
		"role":  "developer",
	}
	// Get value according to specific key
	fmt.Println(user["name"])
	// When trying to get value from non-exist key, it returns empty
	fmt.Println("Age is:", user["age"])
	// Proper way to dealing with the situation when it's possible the key is not existed
	age, ok := user["age"]
	if ok == true {
		fmt.Println("Age:", age)
	} else {
		fmt.Println("Age not found")
	}
	// Add item
	user["age"] = "118"
	fmt.Println(user["age"])
	// Iterate maps
	for key, value := range user {
		fmt.Println(key, value)
	}
}
