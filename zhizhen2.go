package main

import "fmt"

func main()  {
	a, b := 10, 20
	fmt.Println(a, b)
	swap2(&a, &b)
}

func swap2(a, b *int)  {
	*a, *b = *b, *a
	fmt.Println("swap2 a", *a, "b", *b)
}
