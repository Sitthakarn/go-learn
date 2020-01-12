package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@email.com",
			zipCode: 94000,
		},
	}

	dew := person{
		firstname: "Sittha",
		lastname:  "Prasomsup",
		contactInfo: contactInfo{
			email:   "dew@email.com",
			zipCode: 10140,
		},
	}
	dew.print()
	fmt.Println("")
	dew.updateName("Sitthakarn")
	dew.print()
	fmt.Println("")
	fmt.Println("")

	alex.print()
	fmt.Println("")
	alex.updateName("Alexson")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
