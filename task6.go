package main

import (
	"fmt"
	"net/mail"
	"strconv"
)

type Student struct {
	id        string
	firstname string
	lastname  string
}

// Contact {firstname, lastname, phone, email, age}
type Contact struct {
	Student
	phone string
	email string
	age   string
}

func main() {
	var c1 *Contact = new(Contact)

	// +Read Contact
	read(c1)

	fmt.Printf("%+v\n", c1)

	// + Display Contact
	display(*c1)
	// fmt.Printf("%+v\n", c1)
}

func read(c *Contact) {
	boolean := true
	for boolean {
		fmt.Printf("Enter id: ")
		fmt.Scanf("%s", &c.Student.id)
		_, err := strconv.Atoi(c.Student.id)
		if err != nil {
			fmt.Printf("Wrong input, please enter again %v ", err)
		} else {
			boolean = false
		}
	}

	for boolean {
		fmt.Printf("Enter phone: ")
		fmt.Scanf("%s", &c.phone)
		_, err := strconv.Atoi(c.phone)
		if err != nil {
			fmt.Printf("Wrong input, please enter again %v  ", err)
		} else {
			boolean = false
		}
	}

	for boolean {
		fmt.Printf("Enter email: ")
		fmt.Scanf("%s", &c.email)
		_, err := mail.ParseAddress(c.email)
		if err != nil {
			fmt.Printf("Wrong input , please enter again %v ", err)
		} else {
			boolean = false
		}
	}

	for boolean {
		fmt.Printf("Enter age: ")
		fmt.Scanf("%d", &c.age)
		_, err := strconv.Atoi(c.age)
		if err != nil {
			fmt.Printf("Wrong input, please enter again  %v", err)
		} else {
			boolean = false
		}
	}
	fmt.Printf("Enter firstname: ")
	fmt.Scanf("%s", &c.Student.firstname)

	fmt.Printf("Enter lastname: ")
	fmt.Scanf("%s", &c.Student.lastname)

}

func display(c Contact) {
	fmt.Printf("Phone: %s\n", c.Student.id)
	fmt.Printf("Firstname: %s\n", c.Student.firstname)
	fmt.Printf("Lastname: %s\n", c.Student.lastname)
	fmt.Printf("Phone: %s\n", c.phone)
	fmt.Printf("Email: %s\n", c.email)
	fmt.Printf("Age: %s\n", c.age)
}
