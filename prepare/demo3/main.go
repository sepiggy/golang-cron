package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	output []byte
	err    error
}

func main() {
	// 在一个goroutine里执行一个cmd, 让它执行2秒
	// 在1秒的时候，杀死这个cmd

	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	resultChan = make(chan *result)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)

		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2; echo hello;")

		output, err = cmd.CombinedOutput()

		resultChan <- &result{
			output: output,
			err:    err,
		}

		close(resultChan)
	}()

	time.Sleep(time.Second * 1)

	cancelFunc()

	res = <-resultChan

	fmt.Println(string(res.output), res.err)
}
