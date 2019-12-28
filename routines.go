package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {

		go count(i)
	}

	time.Sleep(time.Millisecond * 21000)

}

func count(id int) {

	for i := 0; i < 10; i++ {
		fmt.Println("routine number", id, ":", "number inside count loop", i)

		time.Sleep(time.Millisecond * 1000)

	}

}
