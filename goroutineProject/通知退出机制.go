package goroutineProject

import (
	"fmt"
	"math/rand"
)

func dosth(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done: //done有值后，两个值都为真，所以总会进入这个case停止循环
				break Label
			}
		}
		close(ch)
	}()
	return ch
}
func Tongzhituichu() {
	done := make(chan struct{})
	ch := dosth(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)       //close之后，子函数再取会取到done的零值，所以这里和done<-struct{}{}作用一样
	fmt.Println(<-ch) //由于ch已经关闭，所以取到的是零值
}
