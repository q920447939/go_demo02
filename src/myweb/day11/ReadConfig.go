package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	configer, e := config.NewConfig("ini", "./config.conf")
	if e != nil {
		fmt.Println("not find this file")
	}
	port, _ := configer.Int("service::port")
	fmt.Println(port)
	configMap := make(map[string]interface{})
	configMap["level"] = logs.LevelDebug
	configMap["filename"] = "./logApple.log"
	bytes, _ := json.Marshal(configMap)
	logs.SetLogger(logs.AdapterFile, string(bytes))
	logs.Debug("Debug one ,name equal [%s]", "zhangsan")
	logs.Warn("Warn one ,name equal [%s]", "zhangsan")
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
