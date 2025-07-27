// This file is used to test the const usage in Go

package main

import (
	"fmt"
)


// Function to print weekdays
func printWeekdays() {

	type weekDay int

	const (
		Sunday weekDay = iota + 1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Weekdays:")
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}


// Function to convert bytes to GB (with parameters and return value)
func bytesToGB(bytes float64) float64 {
	const (
		_  = iota             // ignore first value by assigning to a blank identifier
		KB = 1 << (10 * iota) // bit shift 1 to the left by 10 times iota
		MB                    // 1 << (10 * 2) = 1048576
		GB                    // 1 << (10 * 3) = 1073741824
		TB                    // 1 << (10 * 4) = 1099511627776
		PB                    // 1 << (10 * 5) = 1125899906842624
		EB                    // 1 << (10 * 6) = 1152921504606846976
		ZB                    // 1 << (10 * 7) = 1180591620717411303424
		YB                    // 1 << (10 * 8) = 1208925819614629174706176
	)

	return bytes / float64(GB)
}

// Function to print file size in GB
func printFileSize(bytes float64) {
	gb := bytesToGB(bytes)
	fmt.Printf("File size: %.2fGB\n", gb)
	{
		{ /*
				The %.2f is a format specifier used with fmt.Printf to control how floating-point numbers are displayed:
				%f - Tells Printf to format the value as a floating-point number (decimal notation)
				.2 - Specifies the precision - exactly 2 decimal places after the decimal point
				So %.2f means: "display as a float with exactly 2 decimal places"
			*/
		}
	}
}

func main() {
	// Call the weekdays function
	printWeekdays()

	// Call the file size functions
	fileSize := 4000000000.0
	printFileSize(fileSize)

}
