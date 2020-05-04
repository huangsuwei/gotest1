package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "i am in the world i like the world"

	mRet := WordCountFunc(str)

	fmt.Println(mRet)
}

func WordCountFunc(str string) map[string]int {
	m1 := make(map[string]int)
	s1 := strings.Fields(str)

	for i := 0; i < len(s1); i ++ {
		if _, has := m1[s1[i]]; has {
			m1[s1[i]] += 1
		} else {
			m1[s1[i]] = 1
		}
	}

	for _, value := range s1 {
		fmt.Println(value)
	}

	return m1
}