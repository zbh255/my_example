package main

// 编译之后通过参数运行

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	background("/Users/harder/github.com-codes/bups/_build_release/log/service.log")
}

//链接：https://zhuanlan.zhihu.com/p/146192035
func background(logFile string) error {
	//os.Args 是一个切片,保管了命令行参数，第一个是程序名
	//go程序启动时不包含管道符了,就直接运行了
	cmd := exec.Command(os.Args[0], os.Args[1:]...)

	//若有日志文件, 则把子进程的输出导入到日志文件
	if logFile != "" {
		stdout, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println(os.Getpid(), ": 打开日志文件错误:", err)
			return err
		}
		cmd.Stderr = stdout
		cmd.Stdout = stdout
	}

	//异步启动子进程
	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}