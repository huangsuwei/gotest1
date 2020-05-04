package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	go sing()
	go dance()
	for {
		;
	}
}

func sing()  {
	fmt.Println(runtime.GOROOT())
	for i := 0; i < 5; i ++ {
		runtime.Gosched()
		fmt.Println("sing----")
		time.Sleep(100 *time.Millisecond)
	}
}

func dance()  {
	for i := 0; i < 5; i ++ {
		fmt.Println("dancing----")
		time.Sleep(100 *time.Millisecond)
	}
}
