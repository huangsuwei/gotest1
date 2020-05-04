package main

import "fmt"

func main()  {
	ch := make(chan int)
	ch <- 87
	num := <- ch
	fmt.Println("num = ", num)
}
