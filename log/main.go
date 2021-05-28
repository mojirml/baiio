package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
)

func main() {
	uuids := uuid.NewV1()
	log.SetPrefix(uuids.String() + " ")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Print("我就是一条日志")
	log.Printf("%s,", "谁说我是日志了，我是错误")
}

func init() {
	fmt.Println("init main")
}
