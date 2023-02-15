package main

import (
	"fmt"
)

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	gg        address
}

func main() {
	p := person{firstName: "Elon", lastName: "Musk", gg: address{city: "Los Angeles", state: "California"}}
	fmt.Println(p.gg, p.gg.city)

	//p.fullAddress() //accessing fullAddress method of address struct

}
