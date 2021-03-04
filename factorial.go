package main

import "fmt"

func IteractiveFactorial(number int) uint64 {
	var result uint64 = 1
	if number < 0 {

	} else {
		for i := 1; i <= number; i++ {
			result *= uint64(i)
		}
	}
	return result
}
func main() {
	var number int
	fmt.Print("Enter a number : ")
	fmt.Scanln(&number)
	fmt.Printf("The factorial of %d = %d", number, IteractiveFactorial(number))
}
