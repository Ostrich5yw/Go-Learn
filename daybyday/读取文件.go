package daybyday

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadByBlock(filepath string) {
	var fp *os.File
	var err error
	if fp, err = os.OpenFile(filepath, os.O_RDWR, 6); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer fp.Close()

	output := make([]byte, 1024)
	for {
		size, err := fp.Read(output)
		if err == io.EOF {
			return
		}
		fmt.Println(output[:size])
	}

}

func ReadByLine(filepath string) {
	var fp *os.File
	var err error
	if fp, err = os.OpenFile(filepath, os.O_RDWR, 6); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer fp.Close()
	readBuf := bufio.NewReader(fp)
	for {
		line, over, _ := readBuf.ReadLine()
		if !over {
			break
		}
		fmt.Println(line)
	}

	for {
		lineByte, err := readBuf.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		fmt.Println(string(lineByte))
	}
}
