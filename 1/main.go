// Race Condition
package main

import (
	"fmt"
	"time"
)

var data int

func main() {
	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Println("Hello data is ", data)
	}
}
