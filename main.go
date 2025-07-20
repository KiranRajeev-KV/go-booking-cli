package main

import (
	"bufio"
	"fmt"
	"go-cli/helper"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Booking struct {
	username     string
	ticketsToBuy int
}

// package level variable
var bookings []Booking
var bookingsMap []map[string]string

var wg = sync.WaitGroup{}

func main() {

	var welcomeMessage = "learning go by making a go-cli"
	fmt.Println(welcomeMessage)

	//datatypes : uint,int,float,complex,string
	var eventName string = "Code Hack"
	const maxSeats = 100

	// ':=' -> creates a variable but cant cant assgin a specific type
	remainingSeats := 100

	// arrays
	// var bookings = [100]string{}
	// var bookings [50]string

	// slices
	// bookings := []Booking{}

	fmt.Println("Welcome to Event Booking CLI in go")
	fmt.Println("Max seats for any event is", maxSeats)

	// fmt.Println("We have", remainingSeats, "seats available for", eventName)
	// %v -> value of variable
	// %T -> type of variable
	// %t -> bool
	// %d -> int
	// %g -> float
	// %s -> string
	fmt.Printf("We have %v seats available for %v\n", remainingSeats, eventName)

	// for remainingSeats > 0 {} -> for loop can also take conditons
	for {
		// declaring a variable with no initialization
		var username string
		// Variables declared without a corresponding initialization are zero-valued
		var ticketsToBuy int

		// fmt.Print("Enter your username : ")
		// fmt.Scan(&username)

		username = getUsername()
		ticketsToBuy = getNoOfTickets()

		if ticketsToBuy > remainingSeats {
			fmt.Printf("We only have %v tickets remaining.\n", remainingSeats)
			continue
		}

		fmt.Printf("\nTickets Booked!\nusername : %v\nNumber of tickets booked: %v\n", username, ticketsToBuy)

		// newBookings := Booking{
		// 	username: username,
		// 	ticketsToBuy: ticketsToBuy,
		// }
		// bookings = append(bookings, newBookings)

		bookings = append(bookings, Booking{username: username, ticketsToBuy: ticketsToBuy})

		// for maps
		var bookingsData = make(map[string]string)
		bookingsData["username"] = username
		bookingsData["ticketsBought"] = strconv.FormatInt(int64(ticketsToBuy), 10)
		bookingsMap = append(bookingsMap, bookingsData)

		wg.Add(1)
		go sendTicket(ticketsToBuy, username)

		fmt.Println(bookingsMap)

		// fmt.Printf("\nPrint bookings slice : %v\n", bookings)
		var firstNames []string = getFirstNames()
		// fmt.Printf("Type of bookings : %T\n", bookings)
		fmt.Printf("First names of all booked users : %v\n", firstNames)

		remainingSeats = remainingSeats - int(ticketsToBuy)

		if remainingSeats == 0 {
			helper.MyOwnPrint("All seats booked. :)")
			break
		}

		fmt.Printf("\nRemaining seats available : %v\n", remainingSeats)
	}
	wg.Wait()
}

func getUsername() string {
	reader := bufio.NewReader(os.Stdin)
	var username string
	for {
		fmt.Print("Enter your username : ")
		// supports full names like "John Doe" with space
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)

		if len(username) >= 3 {
			break
		} else {
			fmt.Println("Error : Username must contain atleast 3 characters.\nTryAgain.")
			continue
		}
	}
	return username
}

func getNoOfTickets() int {
	var ticketsToBuy int
	for {
		fmt.Print("Enter no of tickets to buy : ")
		fmt.Scan(&ticketsToBuy)

		if ticketsToBuy > 0 {
			break
		} else {
			fmt.Println("Error : Number of tickets must be a positive number.")
			continue
		}
	}
	return ticketsToBuy
}

func getFirstNames() []string {
	var firstNames []string
	for _, data := range bookings {
		names := strings.Fields(data.username)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func sendTicket(ticketsToBuy int, username string) {
	time.Sleep(10 * time.Second)

	var ticketDetails = fmt.Sprintf("%v tickets for %v", ticketsToBuy, username)

	fmt.Println("\n#################")
	fmt.Printf("Sending ticket:\n%v\n", ticketDetails)
	fmt.Println("#################")
	wg.Done()
}
