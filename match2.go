package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "3.14 123.123 .68 haha 1.0 abc 7. ab.3 66.6 123"
	ret, err := regexp.Compile(`^(\d)\.\d+`)
	if err != nil {
		fmt.Println("regexp.compile err", err)
		return
	}

	match := ret.FindAllString(str, -1)
	fmt.Println("match:", match)
}
