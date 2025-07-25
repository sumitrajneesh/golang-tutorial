/**
  Program to print the day of the week using  switch case

**/

//

// Program to print the day of the week using fallthrough in switch

//

// Program to check if the day is a weekend or a weekday

package main

import "fmt"

func main() {
  dayOfWeek := "Sunday"

  switch dayOfWeek {
    case "Saturday", "Sunday":
      fmt.Println("Weekend")

    case "Monday","Tuesday","Wednesday","Thursday","Friday":
      fmt.Println("Weekday")

    default:
      fmt.Println("Invalid day")
  }
}