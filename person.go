package main

import "fmt"

type Person struct {
	name string
	age int
	height int
	sex string
}

type Student struct {
	Person
	source float64
}

func main()  {
	var s1 Student
	s1 = Student{Person{"细秀", 26, 160, "女"}, 99.99}
	fmt.Println(s1)
}