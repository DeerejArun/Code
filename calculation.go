package main

import "fmt"

func Calc(x, y int) (int, int) {
	var out1 = x + y
	var out2 = x - y
	return out1, out2
}

func main() {
	num1 := 10
	num2 := 8
	result1, result2 := Calc(num1, num2)
	fmt.Println(result1, result2)

}
