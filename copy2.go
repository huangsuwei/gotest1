package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[7:]
	s2 := data[:5]
	n := copy(s2, s1)
	fmt.Println("n=", n)
	fmt.Println(s2)
}
