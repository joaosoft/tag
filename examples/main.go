package main

import (
	"fmt"
	"tag"
)

type (
	Person struct {
		Id      int    `person:"id"`
		Name    string `person:"name"`
		Age     int    `person:"age"`
		Book    string `book:"book"`
		Address Address
	}
	Address struct {
		Street     string `address:"street"`
		Number     int    `address:"number"`
		PostalCode string `address:"-"`
	}
)

func main() {
	person := Person{
		Id:   1,
		Name: "joao",
		Age:  30,
		Book: "technology for dummies",
		Address: Address{
			Street:     "apple street",
			Number:     1,
			PostalCode: "1234-567",
		},
	}

	// load person
	values := tag.Load(person, "person", "book")
	fmt.Println("Person:")
	for key, value := range values {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}

	// load address
	values = tag.Load(person, "address")
	fmt.Println("\nAddress:")
	for key, value := range values {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}

	// load nil value
	values = tag.Load(nil, "address")
	fmt.Println("\nAddress:")
	for key, value := range values {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}
}
