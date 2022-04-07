package concurrentProject

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

const (
	SERVER_PROTO = "tcp"
	SERVER_ADDR  = "127.0.0.1:8085"
	DELIMITER    = '\t'
)

func myread(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	readbytes, err := reader.ReadBytes(DELIMITER)
	if err != nil {
		return "", err
	}
	return string(readbytes[:len(readbytes)-1]), nil
}

func mywrite(conn net.Conn, content string) (int, error) {
	var nn int
	var err error
	writer := bufio.NewWriter(conn)
	if nn, err = writer.WriteString(content); err != nil {
		return 0, err
	}
	if err := writer.WriteByte(DELIMITER); err != nil {
		return 0, err
	}
	if err := writer.Flush(); err != nil {
		return 0, err
	}
	return nn + 1, nil
}

func handleConn(id int, conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		str, err := myread(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Server[%d] Connection is closed by [Client Address:%s], Read Over...\n", id, conn.RemoteAddr().String())
			} else {
				fmt.Printf("Server[%d] Read Error[Client Address]:%s\n", id, err)
			}
			break
		}
		fmt.Printf("Server[%d] Receive a Message from [Client Address:%s] (content: %s)\n", id, conn.RemoteAddr().String(), str)
		switch str {
		case "你好":
			str = "hello"
		case "你是谁":
			str = "i am tom riddle"
		case "你多大":
			str = "23 years old"
		case "再见":
			str = "bye bye"
		default:
			str = "i cant understand"
		}
		n, err := mywrite(conn, str)
		if err != nil {
			fmt.Printf("Server[%d] Write Error:%s\n", id, err)
			continue
		}
		fmt.Printf("Server[%d] Sent Respone to [Client Address:%s] (written %d bytes, content: %s)\n", id, conn.RemoteAddr().String(), n, str)
	}
}

func ServerGO(id int) error {
	var listener net.Listener
	var conn net.Conn
	var err error
	if listener, err = net.Listen(SERVER_PROTO, SERVER_ADDR); err != nil {
		fmt.Printf("Server[%d] Server Error:%s\n", id, err)
		return err
	}
	defer listener.Close()
	fmt.Printf("Server[%d] Server ready to listen[Local Address:%s]\n", id, listener.Addr().String())
	for {
		if conn, err = listener.Accept(); err != nil {
			fmt.Printf("Connection Error:%s\n", err)
			return err
		}
		fmt.Printf("Server[%d] Connection is established[Client Address:%s]\n", id, conn.RemoteAddr().String())
		go handleConn(id, conn)
	}
}

func ClientGO(id int) {
	ques := []string{"你好", "你是谁", "你多大", "再见", "这句你不会"}
	var conn net.Conn
	var err error
	if conn, err = net.DialTimeout(SERVER_PROTO, SERVER_ADDR, 2*time.Second); err != nil {
		fmt.Printf("Client[%d] has Dial Error:%s\n", id, err)
		return
	}
	defer conn.Close()
	fmt.Printf("Client[%d]: Connect to [Server Address:%s]\n", id, conn.RemoteAddr().String())
	time.Sleep(200 * time.Millisecond)
	conn.SetDeadline(time.Now().Add(20 * time.Second))
	for _, str := range ques {
		var n int
		var strrep string
		if n, err = mywrite(conn, str); err != nil {
			fmt.Printf("Client[%d] Write Error:%s\n", id, err)
			continue
		}
		fmt.Printf("Client[%d] Sent Request to [Server Address:%s] (written %d bytes, content: %s)\n", id, conn.RemoteAddr().String(), n, str)
		strrep, err = myread(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Client[%d] Connection is closed by Server, Read Over...\n", id)
			} else {
				fmt.Printf("Client[%d] Read Error[Server Address:%s]:%s\n", id, conn.RemoteAddr().String(), err)
			}
			break
		}
		fmt.Printf("Client[%d] Receive a Message from [Server Address:%s] (content: %s)\n", id, conn.RemoteAddr().String(), strrep)
	}
}
