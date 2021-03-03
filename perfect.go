package main

import "fmt"

func main() {
	var perfNum, perfsum int
	perfsum = 0
	fmt.Print("Enter the number to find Perfect = ")
	fmt.Scanln(&perfNum)
	for i := 1; i < perfNum; i++ {
		if perfNum%i == 0 {
			perfsum = perfsum + i
		}
	}
	if perfNum == perfsum {
		fmt.Println(perfNum, "is a prfect number")
	} else {
		fmt.Println(perfNum, "is not a perfect number")
	}

}
