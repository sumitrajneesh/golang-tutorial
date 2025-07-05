// package main

// import "fmt"

// func main() {
//     // Declare an array of 5 integers
//     var numbers [5]int

//     // Initialize an array with values
//     names := [3]string{"Alice", "Bob", "Charlie"}

//     // Access elements by index
//     fmt.Println("First name:", names[0]) // Output: Alice

//     // Modify an element
//     numbers[2] = 10
//     fmt.Println("Modified numbers:", numbers) // Output: [0 0 10 0 0]

//     // Get the length of an array
//     fmt.Println("Number of names:", len(names)) // Output: 3

//     // Iterate through an array using a for loop
//     fmt.Println("All numbers:")
//     for i := 0; i < len(numbers); i++ {
//         fmt.Println(numbers[i])
//     }

//     // Iterate through an array using range
//     fmt.Println("All names with index:")
//     for index, name := range names {
//         fmt.Println(index, name)
//     }
// }