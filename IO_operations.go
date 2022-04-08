// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {

// 	// we will be using variable to store name and dispaly on the terminal in next line
// 	// variable declaration in Go is in following format
// 	// format> var then name of the variable and followed with variable type
// 	var firstName string
// 	var secondName string
// 	var birthYear int

// 	// Println function is used to print the output on the terminal
// 	fmt.Println("********************************************************************************")
// 	fmt.Println("This Program is to demonstrate the Input and output Operations in Go Language ")
// 	fmt.Println("Please Enter your First Name")

// 	// Now we will be using ScanLn to obtain user input
// 	fmt.Scanln(&firstName)
// 	fmt.Println("Please Enter your Last Name")
// 	fmt.Scanln(&secondName)
// 	fmt.Println("Please Enter your Date of birth")
// 	fmt.Scanln(&birthYear)

// 	// print function is display output in same line
// 	fmt.Print("User information is:")
// 	// we will be using library called strconv for converting int value to string value
// 	fmt.Println(firstName + " " + secondName + " is born in " + strconv.Itoa(birthYear))

// 	// performing mathematical operations based on declared variable and input provided by use
// 	const currentYear int = 2022
// 	var age int
// 	age = currentYear - birthYear
// 	fmt.Print(firstName + "'s age is " + strconv.Itoa(age))
// }
