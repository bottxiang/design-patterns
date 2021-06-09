package main

import (
	"fmt"
	"pattern/pool"
	"pattern/singleton"
	"pattern/strategy"
)

func main() {
	//001.singleton
	i := singleton.New()
	i["this"] = "that"
	j := singleton.New()
	fmt.Println(j["this"])

	//002.pool
	p := pool.New(2)
	select {
	case obj := <-*p:
		obj.Do()

		*p <- obj
	default:
		// No more objects left — retry later or fail
		return
	}

	//003.strategy
	add := strategy.Operation{strategy.Addition{}}
	fmt.Println(add.Operate(3, 4))

	multiply := strategy.Operation{strategy.Multiplication{}}
	fmt.Println(multiply.Operate(3, 4))
}
