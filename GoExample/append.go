package main

import "fmt"

func Pingpong(s []int) []int{
	s=append(s,3)
	return s
}

func main()  {
	s := make([]int,0)
	fmt.Println(s)
	s = Pingpong(s)
	fmt.Println(s)
}
