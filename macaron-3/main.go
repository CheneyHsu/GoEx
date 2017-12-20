package main

import (
	"log"
	"os"

	"github.com/go-macaron/macaron"
)

type Logger interface {
	Println(...interface{})
}

var logger = log.New(os.Stdout, "[APP]", 0)

func main() {
	m := macaron.New()
	m.Map(logger)

	m.Get("/", func(l *log.Logger) {
		l.Println("这是全局日志")
	})
	m.Get("/123", myLogger, myLogger2, func(l Logger) {
		l.Println("hello world")
	})
	m.Run()
}

func myLogger(ctx *macaron.Context) {
	var logger = log.New(os.Stdout, "[MyLogger]", 0)
	ctx.MapTo(logger, (*Logger)(nil))
}

type SuperLogger struct {
	*log.Logger
}

func myLogger2(ctx *macaron.Context) {
	var logger = log.New(os.Stdout, "[MyLogger2]", 0)
	sl := &SuperLogger{logger}
	ctx.MapTo(sl, (*Logger)(nil))
}
