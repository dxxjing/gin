package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main () {

	buf, _ := exec.Command("date").Output()
	fmt.Println(string(buf))

	cmd := exec.Command("ls", "-a", "-l")
	//给新进程设置 标准输入 标准输出
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("cmd run err:" + err.Error())
		return
	}
	fmt.Println(os.Getpid(), cmd.Process.Pid)
	fmt.Println(out.String())
}
