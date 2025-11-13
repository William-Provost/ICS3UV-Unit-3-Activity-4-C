/*
 * Author: William Provost
 * Version: 1.0.0
 * Date: 2025-11-13
 * Fileoverview: This program tells you what to pay for a movie.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var ageAsString string = ""
	var ageAsNumber int = 0

	var reader = bufio.NewReader(os.Stdin)

	// input age
	fmt.Print("Please enter your age: ")
	ageAsString, _ = reader.ReadString('\n')
	ageAsString = strings.TrimSpace(ageAsString)
	ageAsNumber, _ = strconv.Atoi(ageAsString)

	// check the age
	if ageAsNumber >= 65 {
		fmt.Println("You are a senior")
		fmt.Println("Please pay $15.00 to see the show")
	} else if ageAsNumber >= 19 {
		fmt.Println("You are an adult")
		fmt.Println("Please pay $25.00 to see the show")
	} else if ageAsNumber >= 12 {
		fmt.Println("You are a teenager")
		fmt.Println("Please pay $20.00 to see the show")
	} else if ageAsNumber >= 6 {
		fmt.Println("You are a child")
		fmt.Println("Please pay $15.00 to see the show")
	} else {
		fmt.Println("You are not in school")
		fmt.Println("You are free to see the show")
	}

	fmt.Println("\nDone.")
}
