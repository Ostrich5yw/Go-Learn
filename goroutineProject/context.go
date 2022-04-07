package goroutineProject

import (
	"context"
	"fmt"
	"time"
)

type selfcontext struct {
	context.Context
}

func work(ctx context.Context, name string) { //一个子任务，参数包含上下文信息与子任务名
	for {
		select {
		case <-ctx.Done(): //从上下文接收到结束信息
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) { //一个子任务，参数包含带值的上下文信息与子任务名
	for {
		select {
		case <-ctx.Done(): //从上下文接收到结束信息
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string) //获取上下文中存储的key值
			fmt.Printf("%s is running and value is %s\n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}

func ContextTest() {
	//使用background构建一个WithCancel类型的上下文
	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1") //在此上下文，开启一个goroutine

	//使用ctxa构建一个WithDeadline类型的上下文
	tm := time.Now().Add(3 * time.Second)
	ctxb, t := context.WithDeadline(ctxa, tm)
	go work(ctxb, "work2") //在此上下文，开启一个goroutine
	fmt.Println(t)

	tt := selfcontext{}
	//使用ctxb构建一个WithValue类型的上下文
	mc := selfcontext{ctxb}
	ctxc := context.WithValue(mc, tt, "i am work3")
	go workWithValue(ctxc, "work3") //在此上下文，开启一个goroutine

	time.Sleep(10 * time.Second) //经过10s，由于ctxc，ctxb都包含Deadline上下文，所以都已经超时停止

	//显示调用wordk1的cancal()方法，结束work1
	cancel()
}
