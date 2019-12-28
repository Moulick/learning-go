package main

import "fmt"

func main()  {

	num1, num2 := next(5)

	fmt.Println(num1, num2)

	fmt.Println(add_all(1,2,3,4,5))
}

func next(number int) (int, int) {

	return number+1, number+2
}

func add_all(val ...int) int {

	final_val := 0

	for _, x := range val {

		final_val += x
	}

	return final_val

}