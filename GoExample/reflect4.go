package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{1, "tom", 22}
	set(&u)
	fmt.Println(u)

}

func set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name1")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}
