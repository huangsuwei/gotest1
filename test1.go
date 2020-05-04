package main

import "fmt"

type People interface {

	Speak(string) string

}

type Student1 struct{}

func (stu *Student1) Speak(think string) (talk string) {

	if think == "bitch" {

		talk = "You are a good boy"

	} else {

		talk = "hi"

	}

	return

}

func main() {

	var peo People = Student1{}

	think := "bitch"

	fmt.Println(peo.Speak(think))

}


