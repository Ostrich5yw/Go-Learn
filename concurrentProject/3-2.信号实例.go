package concurrentProject

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

func SendSignal() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal Error: %s\n", err)
			debug.PrintStack()
		}
	}()

	// ps aux | grep "signal" | grep -v "grep" | grep -v "go run" | awk '{print $2}'
	cmds := []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "signal"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("grep", "-v", "go run"),
		exec.Command("awk", "{print $2}"),
	}
	output, err := RunCmds(cmds) //运行命令，得到当前进程的Pid
	if err != nil {
		fmt.Printf("Command Execution Error: %s\n", err)
		return
	}
	pids, err := GetPids(output) //将其[]string转化为[]int
	if err != nil {
		fmt.Printf("PID Parsing Error: %s\n", err)
		return
	}
	fmt.Printf("Target PID(s):\n%v\n", pids)

	for _, pid := range pids { //向进程发送信号，进行测试
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Process Finding Error: %s\n", err)
			return
		}
		sig := syscall.SIGQUIT
		fmt.Printf("Send signal '%s' to the process (pid=%d)...\n", sig, pid)
		err = proc.Signal(sig)
		if err != nil {
			fmt.Printf("Signal Sending Error: %s\n", err)
			return
		}
	}
}
func HandleSignal() {
	sigRecv1 := make(chan os.Signal, 1) //声明两个自定义处理通道，处理当前进程接收到的信号
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for rec := range sigRecv1 {
			fmt.Printf("接收到来自SigRecv1的信号：%s\n", rec)
		}
		fmt.Printf("End. [sigRecv1]\n")
		wg.Done()
	}()
	go func() {
		for rec := range sigRecv2 {
			fmt.Printf("接收到来自SigRecv2的信号：%s\n", rec)
		}
		fmt.Printf("End. [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("Wait for 2 seconds... ")
	time.Sleep(2 * time.Second)
	fmt.Printf("停止SigRecv2自定义通道...")
	signal.Stop(sigRecv2)
	close(sigRecv2)
	fmt.Printf("done. [sigRecv2]\n")

	wg.Wait()
}
