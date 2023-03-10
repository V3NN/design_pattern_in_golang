package proxy

import "fmt"

// 商品
type Goods struct {
	Name string
	Fact bool
}

// 抽象一个购物主题
type Shopping interface {
	Buy(goods *Goods)
}

// 实现层
type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国购买了", goods.Name)
}

type AmericanShopping struct {
}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国购买了", goods.Name)
}

// 代理
type Proxy struct {
	Shopping Shopping
}

func NewProxy(s Shopping) *Proxy {
	return &Proxy{
		Shopping: s,
	}
}

func (p *Proxy) Check(goods *Goods) bool {
	if goods.Fact != true {
		fmt.Println("商品[", goods.Name, "] 为假，不建议购买")
	}
	return goods.Fact
}

func (p *Proxy) Buy(goods *Goods) {
	if p.Check(goods) == true {
		p.Shopping.Buy(goods)
	}
}
