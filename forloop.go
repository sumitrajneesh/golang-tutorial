package main

import "fmt"

func main() {
	/*
	simple loop
	*/
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Hi go")
	// }

	/*
	 Iinitial statement and the post statement from the for syntax, and only use the condition
	*/

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
}