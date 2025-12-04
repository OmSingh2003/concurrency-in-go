// solutionb ot problem of memory synchronization code in not that grate
package main

import (
	"fmt"
	"sync"
)

var (
	memory sync.Mutex
	value  int = 0
)

func main() {
	go func() {
		memory.Lock()
		value++
		memory.Unlock()
	}()
	// time.Sleep(2 * time.Second)
	memory.Lock()
	if value == 0 {
		fmt.Printf("value is zero \n")
	} else {
		fmt.Printf("value is : %d \n", value)
		memory.Unlock()
	}
}
