// package main

// import (
// 	"fmt"
// 	"strings"
// )

// var grades = map[string]float64{
// 	"S": 10,
// 	"A": 9,
// 	"B": 8,
// 	"C": 7,
// 	"D": 6,
// 	"E": 5,
// 	"F": 0,
// }

// func main() {

// 	// refference change variable nameshttps://stackoverflow.com/questions/59623663/trouble-creating-gpa-calculator

// 	totalCredits := 0.0
// 	totalScore := 0.0

// 	var subjects int
// 	for {
// 		fmt.Println("Enter number of subjects:")
// 		n, err := fmt.Scan(&subjects)
// 		if n == 1 && err == nil {
// 			if subjects >= 0 && subjects <= 20 {
// 				break
// 			}
// 		}
// 	}

// 	for i := 0; i < subjects; i++ {
// 		var grade float64
// 		for {
// 			var letterGrade string
// 			fmt.Println("Enter letter grade:")
// 			n, err := fmt.Scanf("%s", &letterGrade)
// 			if n == 1 && err == nil {
// 				var ok bool
// 				grade, ok = grades[strings.ToUpper(letterGrade)]
// 				if ok {
// 					break
// 				}
// 			}
// 		}

// 		var credits float64
// 		for {
// 			fmt.Println("Enter credits:")
// 			n, err := fmt.Scan(&credits)
// 			if n == 1 && err == nil {
// 				if credits >= 0 && credits <= 10 {
// 					break
// 				}
// 			}
// 		}

// 		totalCredits += credits
// 		totalScore += credits * grade
// 	}

// 	gpa := 0.0
// 	if totalCredits != 0.0 {
// 		gpa = totalScore / totalCredits
// 	}
// 	fmt.Println("GPA:", gpa)
// }
