package main

import "fmt"

type Person1 struct {
	name string
	age int
	hobby []string
}

func main()  {
	p1 := new(Person1)
	SetParam(p1)
	fmt.Println(p1)
}

func SetParam(person *Person1)  {
	person.name = "huangsuwei"
	person.age = 26
	person.hobby = append(person.hobby, "play basketball")
	person.hobby = append(person.hobby, "shopping")
}
