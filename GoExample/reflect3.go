package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)
}
