package singleton

import (
	"sync"
)

var (
	once     sync.Once
	instance Singleton
)

type Singleton map[string]string

func New() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})
	return instance
}
