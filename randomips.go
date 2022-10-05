/*
randomips generates a single, random, IP address.
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func makeIpAddress() string {

	// The random module will return the same number unless it is seeded.  So we use Unix Nano
	// time to do just that
	rand.Seed(time.Now().UnixNano())

	// Create the individual octets of the IP address using the random number generator
	// with a max value of 255
	firstOctet := rand.Intn(255)
	secondOctet := rand.Intn(255)
	thirdOctet := rand.Intn(255)
	fourthOctet := rand.Intn(255)

	// Convert the random numbers from an integer to a string and format the octets into the final
	// IP address
	ipAddress := strconv.Itoa(firstOctet) + "." + strconv.Itoa(secondOctet) + "." + strconv.Itoa(thirdOctet) + "." + strconv.Itoa(fourthOctet)

	return ipAddress
}

func main() {
	fmt.Println(makeIpAddress())
}
