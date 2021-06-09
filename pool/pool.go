package pool

import "fmt"

type Pool chan *Object

type Object struct {}

func (o Object)Do() {
	fmt.Println("Do something...")
}

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}
