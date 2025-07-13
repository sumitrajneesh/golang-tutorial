// package main

// import "fmt"

// func main() {
// 	/*
// 	simple loop
// 	*/
// 	// for i := 0; i < 3; i++ {
// 	// 	fmt.Println("Hi go")
// 	// }

// 	/*
// 	 Iinitial statement and the post statement from the for syntax, and only use the condition
// 	*/

// 	/*
// 	i := 0

// 	for i < 5 {
// 		fmt.Println(i)
// 		i++
// 	}
// 	*/

	
// 	/*
// 	buf := bytes.NewBufferString("one\ntwo\nthree\nfour\n")

// 	for {
// 		line, err := buf.ReadString('\n')
// 		if err != nil {
// 			if err == io.EOF {

// 				fmt.Print(line)
// 				break
// 			}
// 			fmt.Println(err)
// 			break
// 		}
// 		fmt.Print(line)
// 	}
// 	*/

// 	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}
// /*
// 	for i := 0; i < len(sharks); i++ {
// 		fmt.Println(sharks[i])
// 	}
// */

// /*
// 	Using Range clause
// */
// for i, shark := range sharks {
// 		fmt.Println(i, shark)
// 	}

// }