package main

import (
	"fmt"
	"log"
)

var printFn = func(idx int, bean interface{}) error {
	fmt.Printf("%d: %#v\n", idx, bean.(*Account))
	return nil
}

func main() {
	fmt.Println("Welcome to bank of xorm!")

	count, err := getAccountCount()
	if err != nil {
		log.Fatalf("Fail to get account count: %v\n", err)
	}
	fmt.Println("Account count:", count)

	for i := count; i <= 10; i++ {
		if err = newAccount(fmt.Sprintf("joe%d", i), float64(i)*100); err != nil {
			log.Fatalf("Fail to create account: %v\n", err)
		}
	}

	//迭代查询
	fmt.Println("Query all records:")
	x.Iterate(new(Account), printFn)
	fmt.Println()
	a := new(Account)

	rows, err := x.Rows(new(Account))
	if err != nil {
		log.Fatalf("Fail to get rows: %#v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(a); err != nil {
			log.Fatalf("Fail to get row: %v\n", err)
		}
		fmt.Printf("%#v\n", a)
	}
	//查询特定字段
	fmt.Println("\nOnly query name:")
	x.Cols("Name").Iterate(new(Account), printFn)

	//排除特殊字段
	fmt.Println("\nOuery all but name:")
	x.Omit("name").Iterate(new(Account), printFn)

	//查询结果偏移
	fmt.Println("\nOffset 2 and limit 3:")
	x.Limit(3, 2).Iterate(new(Account), printFn)
}
