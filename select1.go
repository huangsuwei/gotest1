package main

import (
	"fmt"
	"runtime"
)

func fbnq(ch <- chan int, quit chan bool)  {
	for {
		select {//case 里面必须是io操作
			case num := <- ch :
				fmt.Println(" ", num)
			case <- quit :
				runtime.Goexit()
		}
	}
}

func main()  {
	ch := make(chan int)
	quit := make(chan bool)

	go fbnq(ch, quit)

	x, y := 1, 1
	for i := 0; i < 20; i ++ {
		ch <- x
		x, y = y, x + y
	}
	quit <- true
}
