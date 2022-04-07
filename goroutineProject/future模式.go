package goroutineProject

import (
	"fmt"
	"time"
)

type query struct {
	in  chan string
	out chan string
}

func doquery(q query) {

	in := <-q.in
	time.Sleep(5 * time.Second) //模拟访问数据库进行查询
	q.out <- in + "的结果"

}

func Futuretest() {
	temp := query{in: make(chan string), out: make(chan string)}
	go doquery(temp)
	temp.in <- "select * from user"

	time.Sleep(3 * time.Second) //模拟在进行数据库查询时，我们可以做一些其他的事情

	fmt.Println(<-temp.out) //做完其他的事情，并且数据库查询也完成了，我们从通道中获取查询结果
}
