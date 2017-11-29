package main

import (
	"fmt"
	"log"

	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件 %s", err)
	}

	log.Printf("作者: %s", cfg.MustValue("", "author"))
	log.Printf("GitHub: %s", cfg.MustValue("", "mygithub"))
	comment := cfg.GetSectionComments("courses")
	log.Printf("%s:", comment)

	for i := 1; i <= 3; i++ {
		log.Printf("#%d:%s", i, cfg.MustValue("courses", fmt.Sprintf("#%d", i)))
	}
	log.Printf("%s", cfg.MustValue("dir.Go名库讲解.01-goconfig", "name"))

	cfg.SetKeyComments("courses", "#3", "https://github.com/Unknwon/go-rock-libraries-showcases")

	if err = goconfig.SaveConfigFile(cfg, "output1.ini"); err != nil {
		log.Fatalf("无法保存配置文件: %v", err)
	}
}
