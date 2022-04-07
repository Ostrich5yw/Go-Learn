package goroutineProject

import "fmt"

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v + 1
		}
		close(out)
	}()
	return out
}

func Guandao() {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()
	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}
}
