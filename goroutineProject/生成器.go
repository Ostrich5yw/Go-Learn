package goroutineProject

import (
	"math/rand"
)

func Shengchengqi() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
