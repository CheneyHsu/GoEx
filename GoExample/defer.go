package main

import "fmt"

func main() {
	a()
	b()
	c()
}

func a() {
	fmt.Println("a")
}
func b() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("revover in b")
		}
	}()
	panic("panic in b")

}

func c() {
	fmt.Println("c")

}
