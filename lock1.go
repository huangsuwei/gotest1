package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex//锁只有一把，两个属性，读和写
var value int//定义一个共享数据区

func main(){
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())

	//quit := make(chan bool)
	ch := make(chan int)

	for i := 0; i < 5; i ++ {
		go goRead(ch, i + 1)
	}
	for i := 0; i < 5; i ++ {
		go goWrite(ch, i + 1)
	}

	//<- quit
	for {
		;
	}
}

func goRead(in <-chan int, idx int)  {
	for {
		//rwMutex.RLock()
		num := <-in
		fmt.Printf("----%dth读go程，读出%d\n", idx, num)
		//rwMutex.RUnlock()
	}
}

func goWrite(out chan<- int, idx int)  {
	for {
		num := rand.Intn(1000)
		rwMutex.Lock()
		out <- num
		fmt.Printf("%dth写go程，写入：%d\n", idx, num)
		time.Sleep(300 *time.Millisecond)
		rwMutex.Unlock()
	}
}


