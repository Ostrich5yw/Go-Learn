package goroutineProject

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	num int
}


   
	
func (t *task) job(wait sync.WaitGroup) {
	fmt.Println("i am", t.num, "WORKING")
	time.Sleep(3 * time.Second) //模拟工作流程
	fmt.Println("i am", t.num, "DONE")
	wait.Done() //工作完成，wait-1
}

func AllocateJob() {
	wait := sync.WaitGroup{}         //wait防止程序未执行完成就退出
	taskchan := make(chan *task, 10) //接收任务通道
	go func() {
		for i := 1; i <= 10; i++ { //模拟有新工作加入，1秒加入一个
			task := &task{num: i}
			wait.Add(1) //wait+1
			time.Sleep(1 * time.Second)
			taskchan <- task //任务加入任务通道
		}
	}()
	go func() {
		for task := range taskchan { //从通道中取出任务，并分配goroutine执行
			go task.job(wait)
		}
	}()
	wait.Wait()
}
