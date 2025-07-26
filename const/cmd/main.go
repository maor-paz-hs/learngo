package main

import (
	"fmt"
)

const (
	_ = iota // ignore first value by assigning to a blank identifier

	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000. // This will result float number
	fmt.Printf("%.2fGB", fileSize/GB) //
}
