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
	var courseArray [2]course
	for i, s := range studentArray {

		fmt.Println(i, s)
		fmt.Println(" Please enter the data for the following student id", s.id)
		courseArray = enterCourseInformation()

		fmt.Println()
		var n = 2
		var j int
		for j = 0; j < n; j++ {
			s.courses[j] = courseArray[j]
			fmt.Println(courseArray)
			print_courseDetails(courseArray[j])
		}
	}

	// for i, s := range studentArray {
	// 	fmt.Println(i, s)
	// 	fmt.Println(" Please enter the data for the following student id", s.id)

	// 	print_courseDetails(s.courses[0])

	// }

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
func findGradeLetter(sObject student) {

	// refference change variable nameshttps://stackoverflow.com/questions/59623663/trouble-creating-gpa-calculator

	totalCredits := 0.0
	totalScore := 0.0

	var subjects int
	for {
		fmt.Println("Enter number of subjects:")
		n, err := fmt.Scan(&subjects)
		if n == 1 && err == nil {
			if subjects >= 0 && subjects <= 20 {
				break
			}
		}
	}

	for i := 0; i < subjects; i++ {
		var grade float64
		for {
			var studentletterGrade string
			fmt.Println("Enter letter grade:")
			n, err := fmt.Scanf("%s", &studentletterGrade)
			if n == 1 && err == nil {
				var ok bool
				grade, ok = grades[strings.ToUpper(studentletterGrade)]
				if ok {
					break
				}
			}
		}

		var credits float64
		for {
			fmt.Println("Enter credits:")
			n, err := fmt.Scan(&credits)
			if n == 1 && err == nil {
				if credits >= 0 && credits <= 10 {
					break
				}
			}
		}

		totalCredits += credits
		totalScore += credits * grade
	}

	gpa := 0.0
	if totalCredits != 0.0 {
		gpa = totalScore / totalCredits
	}
	fmt.Println("GPA:", gpa)

}

// printing the course information
func print_courseDetails(sObject course) {
	fmt.Printf("For the following course name : %s, taught by %s, student has current lettergrade %s and grade in the class is %d %", sObject.courseName, sObject.instructorName, sObject.gradeLetter, sObject.studentgrades)
}

// in the go we declare function like below fun functionName (parameter parameter_type )
func print_StudentInformation(sObject student) {
	fmt.Printf("%s %s whos 800 number is %d , and is student at College %s at the age of %d", sObject.firstName, sObject.lastName, sObject.id, sObject.college, sObject.age)
}

func read_data(file string) []student {
	f, errorVal := os.Open(file)

	if errorVal != nil {
		log.Fatal(errorVal)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var str = ""
	var arraylist []student

	for scanner.Scan() {

		str = scanner.Text()
		// fmt.Println()
		split := strings.Split(str, ",")
		// fmt.Print("vale", split[0])
		arraylist = append(arraylist, student{id: split[0], firstName: split[1], lastName: split[2], birthYear: split[3], college: split[4]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("arayList", arraylist)
	return arraylist
}
