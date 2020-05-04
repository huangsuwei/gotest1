package main

import "fmt"

func test1()  {
	fmt.Println("hello work")
}

func test2(i int)  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		//fmt.Println(recover())
		//recover()
	}()

	var arr [10]int
	arr[i] = 100
}

func test3()  {
	fmt.Println("hello work")
}

func main()  {
	test1()
	test2(2)
	test3()
}
