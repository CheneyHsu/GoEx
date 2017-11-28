package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)    //防止未执行完退出，可以使用chan缓存，设置为10
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool,index int){
	a := 1
	for i := 0 ; i<1000000; i++{
		a+=i
	}
	fmt.Println(index,a)
	c<- true
}

*/ //以上是第一种添加缓存的解决方法。

//以下是第二种添加sync waitgroup的解决方法。

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
