package main

//加入_ 意义为只执行这个包的函数,不导入API.
//这样就解决了又导入,又没有使用的问题.
import (
	"errors"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

//定义
type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatalf("Fail to Create engine: %v", err)
	}
	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v", err)
	}
}

func newAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err

}

func getAccount(id int64) (*Account, error) {
	a := &Account{}
	has, err := x.Id(id).Get(a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account not found")
	}
	return a, nil
}

func makeDeposit(id int64, deposit float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		return nil, err
	}
	a.Balance += deposit
	_, err = x.Update(a)
	return a, err
}

func makeWithdraw(id int64, withdraw float64) (*Account, error) {
	a, err := getAccount(id)
	if err != nil {
		return nil, err
	}
	if a.Balance <= withdraw {
		return nil, errors.New("Not enough balance")
	}
	a.Balance -= withdraw
	_, err = x.Update(a)
	return a, err
}

func makeTransfer(id1, id2 int64, balance float64) error {
	a1, err := getAccount(id1)
	if err != nil {
		return err
	}
	a2, err := getAccount(id2)
	if err != nil {
		return err
	}
	if a1.Balance <= balance {
		return errors.New("Not enough balance")
	}
	a1.Balance -= balance
	a2.Balance += balance

	sess := x.NewSession()
	defer sess.Close()

	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Update(a1); err != nil {
		sess.Rollback()
		return err
	} else if _, err = sess.Update(a2); err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func getAccountsAscId() (as []*Account, err error) {
	err = x.Asc("id").Find(&as)
	return as, err

}

func deleteAccount(id int64) error {
	_, err := x.Delete(&Account{Id: id})
	return err
}
