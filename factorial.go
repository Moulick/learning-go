package main

import "fmt"

func main() {

	fmt.Println(factorial(100))
}

func factorial(num float64) float64 {

	if num == 0 {
		return 1
	}

	return num * factorial(num-1)

}
