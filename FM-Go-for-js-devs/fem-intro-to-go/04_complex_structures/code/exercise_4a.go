package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Println(average(120, 5, 7))
}

// func average(num1, num2, num3 float64) float64 {
// 	total := num1 + num2 + num3
// 	return total / 3
// }
