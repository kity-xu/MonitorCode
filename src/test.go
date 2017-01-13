package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//执行【ls /】并输出返回文本
	f, err := exec.Command("sh", "-c", "ps -aux |grep -v grep| grep mysqlds").Output()
	if err != nil {
		fmt.Println(string(err.Error()))
	}
	fmt.Println(string(f))
}
