// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // following code is used for structuring the object

// type student struct {
// 	firstName string
// 	lastName  string
// 	id        string
// 	college   string
// 	birthYear string
// 	age       int
// 	courses   []course
// }
// type course struct {
// 	courseName     string
// 	gradeLetter    string
// 	grades         float32
// 	instructorName string
// }

// func main() {

// 	var fileName string;

// 	read_data(fileName);
// }

// func enterCourseInformation() [5]course {
// 	fmt.Println("********************************************************************************")
// 	fmt.Println("This Program is to demonstrate the Input and output Operations in Go Language ")

// 	//student object variables
// 	var courseObject course
// 	const n int = 5
// 	var courseName, instructorName string
// 	var grades float32

// 	// initializing the array with n values
// 	var courseArray [n]course
// 	var i int

// 	for i = 0; i < n; i++ {
// 		fmt.Println("*****************************************")
// 		fmt.Println("Course" + strconv.Itoa(i+1))
// 		fmt.Println("Please Enter the course Name")

// 		// Now we will be using ScanLn to obtain user input
// 		fmt.Scanln(&courseName)
// 		fmt.Println("Please Enter your grades")
// 		fmt.Scanln(&grades)
// 		fmt.Println("Please Enter instructor Name")
// 		fmt.Scanln(&instructorName)

// 		// storing in student object to put in the array
// 		courseObject.courseName = courseName
// 		courseObject.grades = grades
// 		courseObject.instructorName = instructorName
// 		courseObject.gradeLetter = findGradeLetter(grades)

// 		//putting array object containing all the information
// 		courseArray[i] = courseObject
// 	}

// 	var j int
// 	for j = 0; j < n; j++ {
// 		fmt.Println("-----------------------------------------------")

// 		fmt.Println("Student" + strconv.Itoa(j+1))
// 		print_courseDetails(courseArray[j])
// 		fmt.Println()
// 	}

// 	return courseArray
// }

// //implement grade and write code accordingly TODO
// func findGradeLetter(grades string) {

// }

// // printing the course information
// func print_courseDetails(sObject course) {
// 	fmt.Printf("For the following course name : %s, taught by %s, student has current lettergrade %s and grade in the class is %d %", sObject.courseName, sObject.instructorName, sObject.gradeLetter, sObject.grades)
// }

// // in the go we declare function like below fun functionName (parameter parameter_type )
// func print_StudentInformation(sObject student) {
// 	fmt.Printf("%s %s whos 800 number is %d , and is student at College %s at the age of %d", sObject.firstName, sObject.lastName, sObject.id, sObject.college, sObject.age)
// }

// func read_data(string file) []student {
// 	f, errorVal := os.Open(file)

// 	if errorVal != nil {
// 		log.Fatal(errorVal)
// 	}

// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	var str = ""
// 	var arraylist []student
// 	for scanner.Scan() {
// 		str = scanner.Text()
// 		fmt.Println()
// 		split := strings.Split(str, ",")
// 		fmt.Print("vale", split[0])
// 		arraylist = append(arraylist, student{id: split[0],firstName: split[1],lastName: split[2],birthYear: split[3],college: split[4]})
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Print("arayList", arraylist)
// 	return arraylist
// }
