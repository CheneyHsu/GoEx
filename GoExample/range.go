package main

import "fmt"

func main() {
	num := []int{2, 3, 4}
	sum := 0
	for _, num := range num {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range num {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "bannana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "gotg" {
		fmt.Println(i, c)
	}
}
