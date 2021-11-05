package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	str, err := execShell("ls -l /|head -n 3")
	fmt.Println(err)
	fmt.Println(str)
}

//@link https://www.zhihu.com/people/zh-five
func execShell(s string) (string, error) {
	//这里是一个小技巧, 以'/bin/bash -c xxx'的方式调用shell命令, 则可以在命令中使用管道符,组合多个命令
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out //把执行命令的标准输出定向到out
	cmd.Stderr = &out //把命令的错误输出定向到out

	//启动一个子进程执行命令,阻塞到子进程结束退出
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), err
}
