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

func (u User) Hello() {
	fmt.Println("Hello world.")

}

func main() {
	u := User{1, "ok", 12}
	Info(&u) //这里如果是指针就是要使用t.kind来进行判断
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("Errot")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%20s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s %v\n", m.Name, m.Type)
	}
}
