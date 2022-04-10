package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// following code is used for structuring the object
//https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go
// issues stuck at value nnot changing

type student struct {
	firstName string
	lastName  string
	id        string
	college   string
	birthYear string
	age       int
	courses   [2]course
	gpa       float64
}
type course struct {
	courseName     string
	gradeLetter    string
	studentgrades  float64
	instructorName string
	credit         float64
}

var grades = map[string]float64{
	"S": 10,
	"A": 9,
	"B": 8,
	"C": 7,
	"D": 6,
	"E": 5,
	"F": 0,
}

func main() {

	var fileName string
	var studentArray []student
	fileName = "data.txt"
	studentArray = read_data(fileName)
	// var courseArray [2]course
	// for i, s := range studentArray {

	// 	fmt.Println(i, s)
	// 	fmt.Println(" Please enter the data for the following student id", s.id)
	// 	courseArray = enterCourseInformation()
	// 	for i, s := range studentArray {
	// 		fmt.Println(i)
	// 		var n = 2
	// 		var j int
	// 		for j = 0; j < n; j++ {
	// 			s.courses[j].courseName = courseArray[j].courseName

	// 			s.courses[j].gradeLetter = courseArray[j].gradeLetter
	// 			s.courses[j].studentgrades = courseArray[j].studentgrades
	// 			s.courses[j].instructorName = courseArray[j].instructorName
	// 			s.courses[j].credit = courseArray[j].credit

	// 			// fmt.Println(courseArray)
	// 			// print_courseDetails(courseArray[j])
	// 			fmt.Println("courname", s.courses[j])
	// 		}

	// 	}

	// }

	// issue is not able to write a pointer so that values can be updated
	for i, s := range studentArray {

		fmt.Println("************************************Program Start*************************************")
		fmt.Println(i)
		// var n = 2
		// var j int
		// for j = 0; j < n; j++ {
		// 	fmt.Println("structured course name:", s.courses[i].courseName)
		// 	print_courseDetails(s.courses[j])
		// }
		print_StudentInformation(s)
		// fmt.Println("structured course name1:", s.courses)

	}

}

func enterCourseInformation() [2]course {
	fmt.Println("********************************************************************************")
	fmt.Println("This Program is to demonstrate the Input and output Operations in Go Language ")

	//student object variables
	var courseObject course
	const n int = 2
	var courseName, instructorName, gradeLetter string
	var studentgrades, credit float64

	// initializing the array with n values
	var courseArray [n]course
	var i int

	for i = 0; i < n; i++ {
		fmt.Println("*****************************************")
		fmt.Println("Course" + strconv.Itoa(i+1))
		fmt.Println("Please Enter the course Name")

		// Now we will be using ScanLn to obtain user input
		fmt.Scanln(&courseName)
		fmt.Println("Please Enter your grades")
		fmt.Scanln(&studentgrades)
		fmt.Println("Please Enter instructor Name")
		fmt.Scanln(&instructorName)
		fmt.Println("Please Enter GradeLetter")
		fmt.Scanln(&gradeLetter)
		fmt.Println("Please Enter Course Credit")
		fmt.Scanln(&credit)
		// storing in student object to put in the array
		courseObject.courseName = courseName
		courseObject.studentgrades = studentgrades
		courseObject.instructorName = instructorName
		courseObject.gradeLetter = gradeLetter
		courseObject.credit = credit

		//putting array object containing all the information
		courseArray[i] = courseObject
	}

	// var j int
	// for j = 0; j < n; j++ {
	// 	fmt.Println("-----------------------------------------------")

	// 	fmt.Println("Student" + strconv.Itoa(j+1))
	// 	print_courseDetails(courseArray[j])
	// 	// fmt.Println()
	// }

	return courseArray
}

//implement grade and write code accordingly TODO
func findGradeLetter(cObject [2]course) float64 {

	// refference change variable nameshttps://stackoverflow.com/questions/59623663/trouble-creating-gpa-calculator

	totalCredits := 0.0
	totalScore := 0.0

	// var subjects int = 2
	var grade float64
	var credits float64
	for i := 0; i < len(cObject); i++ {
		grade = grades[strings.ToUpper(cObject[i].gradeLetter)]
		credits = cObject[i].credit

		totalCredits += credits
		totalScore += credits * grade

	}
	gpa := 0.0
	if totalCredits != 0.0 {
		gpa = totalScore / totalCredits
	}
	fmt.Println("GPA:", gpa)
	return gpa
}

// printing the course information
func print_courseDetails(sObject course) {
	fmt.Printf("For the following course name : %s, taught by %s, student has current lettergrade %s and grade in the class is %f ", sObject.courseName, sObject.instructorName, sObject.gradeLetter, sObject.studentgrades)
}

// in the go we declare function like below fun functionName (parameter parameter_type )
func print_StudentInformation(sObject student) {
	fmt.Println("------------------------Student Information-------------------------------")
	fmt.Printf("%s %s whos 800 number is %d , and is student at College %s at the age of %d", sObject.firstName, sObject.lastName, sObject.id, sObject.college, sObject.age)
	fmt.Print("///Course Information////")
	var n = 2
	var j int
	for j = 0; j < n; j++ {
		fmt.Println("structured course name:", sObject.courses[j].courseName)
		print_courseDetails(sObject.courses[j])
	}

}

func read_data(file string) []student {

	fmt.Println("This Program is to demonstrate the Input and output Operations in Go Language ")
	f, errorVal := os.Open(file)

	if errorVal != nil {
		log.Fatal(errorVal)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var str = ""
	var arraylist []student
	// var studentGPA float64
	for scanner.Scan() {

		str = scanner.Text()
		// fmt.Println()
		split := strings.Split(str, ",")
		// fmt.Print("vale", split[0])

		fmt.Println("********************************************************************************")

		courseArray := enterCourseInformation()
		studentGPA := findGradeLetter(courseArray)
		arraylist = append(arraylist, student{id: split[0], firstName: split[1], lastName: split[2], birthYear: split[3], college: split[4], courses: courseArray, gpa: studentGPA})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// var j int
	// for j = 0; j < n; j++ {
	// 	fmt.Println("-----------------------------------------------")

	// 	fmt.Println("Student" + strconv.Itoa(j+1))
	// 	print_courseDetails(courseArray[j])
	// 	// fmt.Println()
	// }

	fmt.Print("arayList", arraylist)
	return arraylist
}
