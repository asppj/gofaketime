package main

import (
	"time"

	"github.com/asppj/gofaketime"
)

func main() {
	println("系统时间", time.Now().Format(time.RFC3339))
	println("启动补丁>>>")
	faker := gofaketime.NewFakeTime()
	defer faker.Close()
	println("修改后的时间", time.Now().Format(time.RFC3339))
	println("运行完成")
}
