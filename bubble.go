package main

import "fmt"

func BubbleSort(arr [5]int)  {
	for i := 0; i < len(arr); i ++  {
		for j := 0; j < i; j ++ {
			if arr[j] > arr[j + 1] {
				var temp int
				temp = arr[j + 1]
				arr[j + 1] = arr[j]
				arr[j] = temp
			}
		}
	}

	for i := 0; i < len(arr); i ++ {
		fmt.Print(arr[i])
	}
}

func main()  {
	var arr = [5]int{1, 4, 3, 5, 2}
	BubbleSort(arr)
}
