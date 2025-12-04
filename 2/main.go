// Memory Access Synchronization
package main

import "fmt"

var data int = 0

func main() {
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("Number is zero")
	} else {
		fmt.Printf("Number is : %d", data)
	}
}
