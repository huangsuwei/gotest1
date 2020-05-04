package main

import "fmt"

func main()  {
	var h int
	var t int
	var u int
	for i := 100; i < 999; i ++  {
		h = i / 100
		t = i % 100  / 10
		u = i % 10
		if i == h * h * h + t * t * t + u * u * u {
			fmt.Println(i)
		}
	}
}
