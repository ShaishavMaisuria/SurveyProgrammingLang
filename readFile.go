// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// type listConst struct {
// 	name string
// }

// func main() {

// 	f, errorVal := os.Open("data.txt")

// 	if errorVal != nil {
// 		log.Fatal(errorVal)
// 	}

// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	var str = ""
// 	var arraylist []listConst
// 	for scanner.Scan() {
// 		str = scanner.Text()
// 		fmt.Println()
// 		split := strings.Split(str, ",")
// 		fmt.Print("vale", split[0])
// 		arraylist = append(arraylist, listConst{name: split[0]})
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Print("arayList",arraylist);
// }
