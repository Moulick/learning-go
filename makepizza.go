package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func main() {

	stringChan := make(chan string)

	for i := 0; i < 3; i++ {

		go makeDough(stringChan)
		go saucing(stringChan)
		go toppingAndBake(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}
}
func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough for", pizzaName, "and send for sauce")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 100)
}

func saucing(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Sauced", pizza, "and sending for topping")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func toppingAndBake(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Toppings added to", pizza, "and baked")

	time.Sleep(time.Millisecond * 100)
}
