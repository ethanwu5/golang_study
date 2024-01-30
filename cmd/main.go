package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// golang  里面执行cmd 命令
func main() {
	cmd := exec.Command("ls", "/")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	ss := strings.Split(string(out), "\n")
	for _, v := range ss {
		fmt.Println(v)
	}
}
