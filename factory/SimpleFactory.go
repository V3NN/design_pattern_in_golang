package simple

import "fmt"

// 基础类
type Person interface {
	Work()
}

// 老师继承了基础类并实现了work方法
type Teacher struct {
	Person
}

func (t *Teacher) Work() {
	fmt.Println("教书")
}

// 司机继承了基础类并实现了work方法
type Driver struct {
	Person
}

func (d *Driver) Work() {
	fmt.Println("开车")
}

// 生产类
type Factory struct {
}

func (f *Factory) CreatePerson(kind string) Person {
	var p Person
	if kind == "teacher" {
		p = new(Teacher)
	} else if kind == "driver" {
		p = new(Driver)
	}
	return p
}

// 业务逻辑层
func main() {
	//创建一个工厂类
	f := new(Factory)

	// 创建一个老师
	t := f.CreatePerson("teacher")
	t.Work()

	// 创建一个司机
	d := f.CreatePerson("driver")
	d.Work()
}
