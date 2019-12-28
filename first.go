package main

import "fmt"

func main() {

	listNum := []int{1, 2, 3, 4, 5, 5, 6, 6, 6}

	fmt.Println("sum is", add(listNum))
}

func add(num []int) float64 {

	x := 0

	for _, val := range num {

		x += val

	}
	return float64(x)
}
