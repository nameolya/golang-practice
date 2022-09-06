package main

import (
	"fmt"
	"strings"
)

func main() {
    conferenceName := "Go Conference"
    const conferenceTickets = 50
    var remainingTickets uint = 50
    bookings := []string{}

    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)

    for {
    var firstName string
    var lastName string
    var email string
    var userTickets uint
 
    fmt.Println("Enter your first name:")
    fmt.Scan(&firstName)
 
    fmt.Println("Enter your last name:")
    fmt.Scan(&lastName)
 
    fmt.Println("Enter your email address:")
    fmt.Scan(&email)
 
    fmt.Println("Enter your number of tickets:")
    fmt.Scan(&userTickets)

    if userTickets <= remainingTickets {
      remainingTickets = remainingTickets - userTickets
      bookings = append(bookings, firstName + " " + lastName)
   
      firstNames := []string{} 
      for _, booking := range bookings {
        var names = strings.Fields(booking) //splits strings with a white space as a separator and returns a slice of strings
        firstNames = append(firstNames, names[0])
      }
   
      fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName,lastName, userTickets, email)
      fmt.Printf("The remaining number of tickets for the %v is %v.\n", conferenceName, remainingTickets) 
      fmt.Printf("These are the first names of the bookings: %v.\n", firstNames)
   
      if remainingTickets  == 0 {
       fmt.Println("Our conference is sold out, please return next year")
       break
      }
    } else {
      fmt.Printf("We only have %v tickets, so you can't book %v tickets.\n", remainingTickets, userTickets)
    }
    
  }
  
}