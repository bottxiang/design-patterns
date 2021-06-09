package pool
/**
* 对象池创建设计模式用于根据需求预期准备和保留多个实例。
*/
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
