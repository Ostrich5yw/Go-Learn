package concurrentProject

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func Test1() {
	cmd0 := exec.Command("echo", "-n", "My first Linux Command By Go")

	//stdout是获取的该命令的Pipe
	stdout, _ := cmd0.StdoutPipe()

	//运行并在出错时报错
	if error := cmd0.Start(); error != nil {
		fmt.Printf("have a problem:%s\n", error)
	}

	//声明一个缓冲区，适用于输出特别多，无法一次通过固定大小的byte[]读完
	var outBuf bytes.Buffer
	for {
		//将管道内容输出到byte[]
		outChannel := make([]byte, 5)
		byteNum, terr := stdout.Read(outChannel)
		if terr != nil {
			if terr == io.EOF { //如果通道中已经没有其他数据，error类型会是io.EOF
				break
			} else {
				fmt.Printf("Error:%s\n", terr)
				return
			}
		}
		if byteNum > 0 {
			outBuf.Write(outChannel[:byteNum])
		}
	}
	//也可以直接用一个缓冲区读取数据
	outBuf2 := bufio.NewReader(stdout)
	output0, _, _ := outBuf2.ReadLine()
	fmt.Printf("the back is:%s\n", outBuf.String())
	fmt.Printf("the back is:%s\n", string(output0))
}
