package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email    string
	mobileNo string
}

func main() {
	//p := person{"Dilanka", "Nimsara"}
	//p1 := person{firstName: "Dilanka", lastName: "Nimsara"}
	//fmt.Println(p)
	//fmt.Println(p1)

	personOne := person{
		firstName: "Dilanka",
		lastName:  "Nimsara",
		contactInfo: contactInfo{
			email:    "dilankanimsara101@gmail.com",
			mobileNo: "0715696661",
		},
	}
	contactInfo1 := contactInfo{
		email:    "dilankanimsara101@gmail.com",
		mobileNo: "0715696661",
	}
	personOne.print()
	contactInfo1.print()

	personOne.updateName("username")
	personOne.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (c contactInfo) print() {
	fmt.Printf("\n %+v", c)
}

func (p person) print() {
	fmt.Printf("\n %+v", p)
}
