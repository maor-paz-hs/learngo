package main

import (
	"fmt"
)

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

func main() {
	fileSize := 4000000000. // This will result float number
	fmt.Printf("%.2fGB", fileSize/GB) // 3.73GB
}
