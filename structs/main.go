package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// In this case, this is equivalent to contactInfo contactInfo
	//contactInfo
}

func main() {
	// Below we are relying on the order of the fields in the struct
	// alex := person{"Alex", "Anderson"}
	// This is much better as we are being explicit about which field we are setting
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}
	// In this case, a 0 value is assigned to the instance of the struct
	var ryan person

	fmt.Println(alex)
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)

	fmt.Println(ryan)

	// Note that we can edit values directly
	ryan.firstName = "Ryan"
	ryan.lastName = "Smith"

	// %+v will print out all the field names and their values
	fmt.Printf("%+v\n", alex)
	fmt.Printf("%+v\n", ryan)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v\n", jim)

	// Using the receiver function
	jim.print()

	// Using the update function
	// jim.updateName("Jimmy")
	// jim.print()
	// We are not writing this
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// The above is not needed as go will automatically pass the pointer adress since the receiver expects a pointer
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Need to use pointers here or the updates will not take hold!
// Pass by value by default
// Need to pass by reference explicitly using pointers so that this function works as expected
//
//	func (p person) updateName(newFirstName string) {
//		p.firstName = newFirstName
//	}
func (pointerToPerson *person) updateName(newFirstName string) {
	// This explicit * is not needed. The dot operator will automatically dereference the pointer
	// (*pointerToPerson).firstName = newFirstName
	pointerToPerson.firstName = newFirstName
}
