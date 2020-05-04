package main

import (
	"fmt"
)

func test(a int, b int) (result int, err error)  {
	err = nil
	/*if b == 0 {
		err = errors.New("被除数不能为0！")
	} else {
		result = a / b
	}*/

	result = a / b

	return result, err
}

func main()  {
	result, err := test(3, 0)
	/*if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println(result)
	}*/

	fmt.Println(result, err)
}
