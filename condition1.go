package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func main()  {
	rand.Seed(time.Now().UnixNano())

	product := make(chan int, 3)
	quit := make(chan bool)
	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i ++ {
		go producer(product, i + 1)
	}

	for i := 0; i < 5; i ++ {
		go customer(product, i + 1)
	}

	<- quit
}

func producer(out chan<- int, idx int)  {
	for {
		cond.L.Lock()
		for len(out) == cap(out) {//这里用for，不用if，if会造成死锁，当条件满足后会执行下面的操作，因为多个并行的时候，条件满足执行大括号，然后下面写数据，写不进去了
			cond.Wait()
		}
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%dth生产者，产生数据%d，公共区域剩余%d个数据\n", idx, num, len(out))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(1 *time.Second)
	}
}

func customer(in <-chan int, idx int)  {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}
		num := <- in
		fmt.Printf("%dth消费者，消费数据%d，公共区域剩余%d个数据\n", idx, num, len(in))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(500 *time.Millisecond)
	}
}