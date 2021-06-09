package main

import (
	"fmt"
	"pattern/singleton"
	"pattern/pool"
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
	case obj := <- *p:
		obj.Do()

		*p <- obj
	default:
		// No more objects left — retry later or fail
		return
	}
}
