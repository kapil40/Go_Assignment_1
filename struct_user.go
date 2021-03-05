package main

import "fmt"

type Address struct {
	city    string
	country string
}

type User struct {
	name    string
	age     int
	address Address
}

func main() {

	p := User{
		name: "Kapil Thakkar",
		age:  22,
		address: Address{
			city:    "Mumbai",
			country: "India",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("Country:", p.address.country)
}
