package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// This type is relate to person struct
type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	// The %+v is used to print a property and a value
	fmt.Printf("%+v", p)
}

// * is give the value that memory address is pointing at
func (p *person) updateName(value string) {
	(*p).firstName = value
}

func main() {
	id := 1
	var zal person
	// There's a two different ways to assign value for the struct
	if id == 1 {
		zal = person{
			firstName: "Farizal",
			lastName:  "Hamami",
			contact: contactInfo{
				email:   "farizalhamami@gmail.com",
				zipCode: 11260,
			},
		}
	}
	if id == 2 {
		zal.firstName = "Farizal"
		zal.lastName = "Hamami"
		zal.contact = contactInfo{
			email:   "farizalhamami@gmail.com",
			zipCode: 11260,
		}
	}
	// & is give the memory address of the value that variable is pointing at
	zalPointer := &zal
	zalPointer.updateName("Zall")
	zal.print()
}
