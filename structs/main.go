package main

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jones",
		contact: contactInfo{
			email: "Jimbo.Jones@email.com",
			zip:   78549,
		},
	}

}
