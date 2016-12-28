package main

import (
	"fmt"
	"os/exec"
)

func shellCommand(paras ...string) string {
	//n := len(paras)
	cmd := exec.Command("python", "haina.im/monitor/monitor_node/pythons/application/mysqld.py", paras[0], paras[1])
	if out, err := cmd.CombinedOutput(); err == nil {
		fmt.Printf("---------%s:", string(out)) //out return value of call *.py
		return string(out)
	} else {
		fmt.Println("running shellCommand error ....")
		return ""
	}
}

func main() {
	shellCommand("nihao", "hello", "xulang", "123")
}
