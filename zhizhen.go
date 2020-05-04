package main

import "fmt"

func main()  {
	var a int = 9

	var p *int = &a

	a = 100

	fmt.Println("a = ", a)

	*p = 250

	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)

	a = 1000
	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)
}
