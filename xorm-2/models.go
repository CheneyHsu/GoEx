package main

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type Account struct {
	Id      float64
	Name    string
	Balance float64
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}
	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

	//记录日志
	//	f, err := os.Create("sql.log")
	//	if err != nil {
	//		log.Fatalf("Fail to create log file:  %v\n", err)
	//	}
	//	x.Logger = xorm.NewSimpleLogger(f)
	//	x.ShowSQL = true

	//设置默认LRU缓存
}

func getAccountCount() (int64, error) {
	return x.Count(new(Account))
}

func newAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err

}
