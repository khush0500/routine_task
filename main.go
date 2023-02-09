package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	fmt.Println("Hello, 世界")
	Working()
}

func Working() {

	c := make(chan int)

	for i := 0; i < 5; i++ {
		go DOSomething(i, c)
	}

	for i := 0; i < 5; i++ {
		mu.Lock()
		fmt.Println(<-c)
		mu.Unlock()

	}
}

func DOSomething(times int, c chan int) {
	time.Sleep(1 * time.Second)
	c <- times

}
