package main

import (
	"log"

	"github.com/Unknwon/goconfig"
)

func main() {

	cfg, err := goconfig.LoadConfigFile("config.ini", "config2.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件: %s", err)
	}
	value, err := cfg.GetValue("", "key_default")
	if err != nil {
		log.Printf("无法获取键值(%s):%s", "key_default", err)
	}
	log.Printf("%s > %s : %s", goconfig.DEFAULT_SECTION, "key_default", value)
	err = cfg.AppendFiles("config3.ini")
	if err != nil {
		log.Fatalf("无法追加配置文件: %s", err)
	}
	vBool := cfg.MustBool("must", "bool404", true)
	log.Printf("%s > %s : %v", "must", "bool404", vBool)
	value, err = cfg.GetValue("", "search")
	if err != nil {
		log.Printf("无法获取键值(%s):%s", "search", err)
	}
	log.Printf("%s > %s : %s", goconfig.DEFAULT_SECTION, "search", value)
	//子分区如果不存在的话,会到父分区进行查找,如果父分区也没有的话,就会报出错误.
	value, err = cfg.GetValue("parent.child", "age")
	if err != nil {
		log.Printf("无法获取键值(%s):%s", "age", err)
	}
	log.Printf("%s > %s : %s", "parent.child", "age", value)

	value, err = cfg.GetValue("parent.child", "sex")
	if err != nil {
		log.Printf("无法获取键值(%s):%s", "sex", err)
	}
	log.Printf("%s > %s : %s", "parent.child", "sex", value)

	sec, err := cfg.GetSection("auto increment")
	if err != nil {
		log.Printf("无法获取分区: %s", err)
	}
	log.Println(sec)

}
