package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	cmd = exec.Command("/bin/bash", "-c", "sleep 2; ls -l; echo hello")

	// 执行命令，捕获子进程的输出
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
