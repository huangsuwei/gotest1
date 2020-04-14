package main

import "fmt"

func main() {
	var h int
	var t int
	var s int
	for i := 100; i <= 999; i++ {
		h = i / 100
		t = i % 100 / 10
		s = i % 10
		if h*h*h+t*t*t+s*s*s == i {
			fmt.Println(i)
		}
	}
}
