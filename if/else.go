// package main

// import "fmt"

// func main() {
//  var age = 20
//  if age >= 18 {
//  fmt.Println("You are an adult")
//  } else {
//  fmt.Println("You are still young")
//  }
// }

// Program to relate two integers using =, > or < symbol

// package main

// import "fmt"

// func main() {

//   number1 := 12
//   number2 := 20

//   if number1 == number2 {
//     fmt.Printf("Result: %d == %d", number1, number2)
//   } else if number1 > number2 {
//     fmt.Printf("Result: %d > %d", number1, number2)
//   } else {
//     fmt.Printf("Result: %d < %d", number1, number2)
//   }
// }

package main

// import "fmt"

// func main() {

//   number1 := 12
//   number2 := 20

//   // outer if statement
//   if number1 >= number2 {

//   // inner if statement
//   if number1 == number2 {
//     fmt.Printf("Result: %d == %d", number1, number2)
//     // inner else statement
//   } else {
//     fmt.Printf("Result: %d > %d", number1, number2)
//   }

//   // outer else statement
//   } else {
//     fmt.Printf("Result: %d < %d", number1, number2)
//   }
// }