package concurrentProject

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func GetPids(strs []string) ([]int, error) { //将Runcmds得到的字符串Pid输出转化为int类型的Pid
	var pids []int             //因为我们运行的是ps aux | grep "signal" | grep -v "grep" | grep -v "go run" | awk '{print $2}'
	for _, str := range strs { //这样一个查询Pid的命令
		pid, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}
	return pids, nil
}

func RunCmds(cmd []*exec.Cmd) (s []string, e error) { //运行一组cmd命令
	var outBuf bytes.Buffer
	var inBuf bytes.Buffer
	var output []byte
	for index, cmd0 := range cmd { //这组cmd命令之间用管道连接，依次连接他们的输入输出
		fmt.Printf("Run Commands:%s\n", cmd0.Path)
		if index != 0 {
			inBuf.Write(output)
			cmd0.Stdin = &inBuf
		}
		cmd0.Stdout = &outBuf
		if err := cmd0.Start(); err != nil {
			return nil, GetError(err, cmd0)
		}
		if err := cmd0.Wait(); err != nil {
			return nil, GetError(err, cmd0)
		}
		output = outBuf.Bytes()

		inBuf.Reset()
		outBuf.Reset()
	}
	outBuf.Write(output) //最后一个cmd的输出，我们按行将其封装为一个[]string
	for {
		line, err := outBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, GetError(err, nil)
			}
		}
		s = append(s, string(line))
	}
	return
}

func GetError(err error, cmd *exec.Cmd, otherstring ...string) error {
	var errMsg string
	if cmd != nil {
		errMsg = fmt.Sprintf("%s [%s %v]", err, (*cmd).Path, (*cmd).Args)
	} else {
		errMsg = fmt.Sprintf("%s", err)
	}
	if len(otherstring) > 0 {
		errMsg = fmt.Sprintf("%s (%v)", errMsg, otherstring)
	}
	return errors.New(errMsg)
}
