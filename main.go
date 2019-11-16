package main

import "fmt"

func main() {

	var twoDArr [][]int

	twoDArr = make([][]int, 3)
	/* for i := 0; i <= 2; i++ {
		twoDArr[i] = make([]int, 3)
	}*/

	a0 := [...]int{1, 2, 3}
	a1 := [...]int{4, 5, 6}
	a2 := [...]int{7, 8, 9}

	twoDArr[0] = a0[:]
	twoDArr[1] = a1[:]
	twoDArr[2] = a2[:]

	for i := 0; i <= 2; i++ {
		fmt.Print(i, ": ")
		for j := 0; j <= 2; j++ {
			fmt.Print(twoDArr[i][j], ", ")
		}
		fmt.Print("\n")
	}
}
