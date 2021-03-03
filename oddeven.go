package main

import "fmt"

func main() {
	fmt.Print("Enter a number : ")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println("It is Even")
	} else {
		fmt.Println("It is Odd")
	}

}
