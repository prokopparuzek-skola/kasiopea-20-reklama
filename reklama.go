package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var S, B, R int
		fmt.Scan(&S, &B, &R)
		if S-R > B {
			fmt.Println("REKLAMU")
		} else {
			fmt.Println("NE REKLAMU")
		}
	}
}
