package lazy

import (
	"fmt"
	"sync"
)

/*
 * 单例三要素：
 * 一是某个类只能有一个实例；
 * 二是它必须自行创建这个实例；
 * 三是它必须自行向整个系统提供这个实例。
 */

var (
	instance *singelton
	once     sync.Once
)

// 保证这个类非公有化，外界不能通过这个类直接创建一个对象
// 那么这个类就应该变得非公有访问 类名称首字母要小写
type singelton struct {
}

func (s *singelton) DoSomeThing() {
	fmt.Println("单例对象的某方法")
}

func GetSingelton() *singelton {
	// sync.once保证只执行一次
	once.Do(func() {
		instance = new(singelton)
	})

	return instance
}

func main() {
	s := GetSingelton()
	s.DoSomeThing()
}
