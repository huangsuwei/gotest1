package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("time:", time.Now())
	myTimer := time.NewTimer(2 *time.Second)
	nowTime := <- myTimer.C
	fmt.Println("now time :", nowTime)
}
