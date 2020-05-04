package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 8; i ++ {
			fmt.Println("son write", i)
			ch <- i
		}
		close(ch)
	}()

	time.Sleep(2 *time.Second)
	for {
		if num, ok := <- ch; ok {
			fmt.Println(num)
		} else {
			fmt.Println("over")
			break
		}
	}
	//for i := 0; i < 8; i ++ {
	//	num := <- ch
	//	fmt.Println("father read", num)
	//}
}
