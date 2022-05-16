package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// Create a struct to match the format of JSON
	type Person struct {
		Name string
		City string
	}

	p := Person{Name: "Marcus", City: "San Francisco"}

	json_data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}

	fmt.Println(string(json_data))

	json_str := `{ "Name": "Marcus", "City": "San Jose"}`
	var person Person

	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	// output result
	fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)

	var people []Person

	file, err := ioutil.ReadFile("names.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	// the names.json file has an array of person objects,
	// so read into people
	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	// output result
	fmt.Println(people)

	/*
	   The JSON field names may not always match the struct names.
	   Go will automatically manage uppercase and lowercase fields,
	   but if they are completely different you can use struct field tags to label.
	*/
	type Person2 struct {
		Name string `json:"username"`
		City string `json:"municipality"`
	}

	person2 := Person2{Name: "Marcus", City: "San Francisco"}
	json_data2, _ := json.Marshal(person2)
	fmt.Println(string(json_data2))

	// You can ignore a field from being encoded or decoded to JSON,
	// using - definition.
	type Person3 struct {
		Name string
		City string
		// The Phone field will be omitted from JSON operations.
		Phone string `json:"-"`
	}

	person3 := Person3{Name: "Marcus3", City: "San Francisco", Phone: "12345"}
	json_data3, _ := json.Marshal(person3)
	fmt.Println(string(json_data3))

	// Set a field to be ignored when empty,
	// use omitempty so JSON encoding will not include the field.
	type Person4 struct {
		Name  string
		City  string
		Phone string `json:",omitempty"`
	}

	person4 := Person4{Name: "Marcus4", City: "San Francisco", Phone: "12345"}
	json_data4, _ := json.Marshal(person4)
	fmt.Println(string(json_data4))

	person4_2 := Person4{Name: "Marcus4_2", City: "San Francisco"}
	json_data4_2, _ := json.Marshal(person4_2)
	fmt.Println(string(json_data4_2))
}
