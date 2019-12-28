// https://www.hackerearth.com/problem/algorithm/the-almond-problem-1ca18502/
package main

import (
	"fmt"
)

func main() {

	var testCases int
	_, _ = fmt.Scanln(&testCases)

	for i := 0; i < testCases; i++ {

		var almondsEat, almondsAdded, almondsInitial int
		_, _ = fmt.Scanln(&almondsEat)
		_, _ = fmt.Scanln(&almondsAdded)
		_, _ = fmt.Scanln(&almondsInitial)

		diff := almondsAdded - almondsEat

		if diff >= 0 {
			if almondsInitial < almondsEat {
				fmt.Println(0)
			} else {
				fmt.Println(-1)
			}
		} else {
			days := (almondsEat-almondsInitial)/(diff*1.0) + 1
			fmt.Println(Max(days, 0))
		}
	}

}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
