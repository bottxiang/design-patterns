package singleton
/**
* 单例创建设计模式将类型的实例化限制为单个对象。
*/
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
