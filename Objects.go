package main

import (
	"fmt"
	"strconv"
)

// following code is used for structuring the object

type student struct {
	firstName string
	lastName  string
	id        int
	college   string
	age       int
}

func main() {
	fmt.Println("********************************************************************************")
	fmt.Println("This Program is to demonstrate the Input and output Operations in Go Language ")

	//student object variables
	var studentObject student
	const n int = 2
	var firstName, lastName, college string
	const currentYear int = 2022

	var birthYear, id, age int
	// initializing the array with n values
	var studentArray [n]student
	var i int

	for i = 0; i < n; i++ {
		fmt.Println("*****************************************")
		fmt.Println("Student" + strconv.Itoa(i+1))
		fmt.Println("Please Enter your First Name")

		// Now we will be using ScanLn to obtain user input
		fmt.Scanln(&firstName)
		fmt.Println("Please Enter your Last Name")
		fmt.Scanln(&lastName)
		fmt.Println("Please Enter your Date of birth")
		fmt.Scanln(&birthYear)
		fmt.Println("Please Enter your ID")
		fmt.Scanln(&id)
		fmt.Println("Please Enter your College Name")
		fmt.Scanln(&college)

		// storing in student object to put in the array
		studentObject.firstName = firstName
		studentObject.lastName = lastName
		age = currentYear - birthYear
		studentObject.age = age
		studentObject.id = id
		studentObject.college = college

		//putting array object containing all the information
		studentArray[i] = studentObject
	}

	var j int
	for j = 0; j < n; j++ {
		fmt.Println("-----------------------------------------------")

		fmt.Println("Student" + strconv.Itoa(j+1))
		print_StudentInformation(studentArray[j])
		fmt.Println()
	}

}

// in the go we declare function like below fun functionName (parameter parameter_type )
func print_StudentInformation(sObject student) {
	fmt.Printf("%s %s whos 800 number is %d , and is student at College %s at the age of %d", sObject.firstName, sObject.lastName, sObject.id, sObject.college, sObject.age)
}
