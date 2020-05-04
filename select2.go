package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
				case num := <- ch :
					fmt.Println("num = ", num)
				case <- time.After(5 * time.Second) :
					quit <- true
					goto label
			}
		}
		label :
			fmt.Println("lable break")
	}()
	for i := 0; i < 3; i ++ {
		time.Sleep(2 * time.Second)
		ch <- i
	}

	<- quit
	fmt.Println("over")
}
