package main

import "fmt"

func mirrorReflection(p int, q int) string {
	var result string
	if q == 0 {
		result = "0"
	} else {
		if q == p {
			result = "1"
		} else {
			a := p / q
			if a % 2 == 0 {
				result = "2"
			} else if a % 3 == 0 {
				result = "1"
			} else {
				result = "不经过任何一个镜子"
			}
		}
	}

	return result
}

func main()  {
	fmt.Println(mirrorReflection(100, 1))
}
