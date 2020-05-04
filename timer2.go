package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("now time", time.Now())
	myTicker := time.NewTicker(time.Second)
	quit := make(chan bool)
	i := 0
	go func() {
		for {
			nowTime := <- myTicker.C
			i ++
			fmt.Println("go now time", nowTime)
			if i == 8 {
				quit <- true
			}
		}
	}()

	<- quit
}
