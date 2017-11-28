package main

import "fmt"

type A struct {
	num int
}

func main() {
	a := A{}
	fmt.Println(a.num)
	a.Incresase()
	fmt.Println(a.num)
}

func (a *A) Incresase() {
	a.num += 100
	fmt.Println("called")

}
