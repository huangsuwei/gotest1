package main

import (
	"fmt"
	"sync"
	"time"
)

var channel = make(chan int)
var mutex sync.Mutex

func printer(s string)  {
	mutex.Lock()
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 *time.Millisecond)
	}
	mutex.Unlock()
}

func person1()  {
	printer("hello")
	channel <- 8
}

func person2()  {
	<- channel
	printer("world")
}

func main()  {
	go person1()
	go person2()
	for {
		;
	}
}
