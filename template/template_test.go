package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	//1. 制作一杯咖啡
	makeCoffee := NewMakeCaffee()
	makeCoffee.MakeBeverage() //调用固定模板方法

	fmt.Println("----------------")

	//2. 制作茶
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}
