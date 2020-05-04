package main

import "fmt"

func main()  {
	channel := make(chan int)

	go func() {
		for i := 0; i < 5; i ++ {
			fmt.Println("sun write", i)
			channel <- i
		}
	}()

	for i := 0; i < 5; i ++ {
		num := <- channel
		fmt.Println("father write", num)
	}
}
