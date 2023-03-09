package hungry

import "fmt"

var instance *singelton

// 保证这个类非公有化，外界不能通过这个类直接创建一个对象
// 那么这个类就应该变得非公有访问 类名称首字母要小写
type singelton struct {
}

func (s *singelton) DoSomeThing() {
	fmt.Println("单例对象的某方法")
}

func Init() {
	instance = new(singelton)
}

func GetSingelton() *singelton {
	return instance
}

func main() {
	instance.DoSomeThing()
}
