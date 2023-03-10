package proxy

import "testing"

func TestProxy(t *testing.T) {
	g1 := &Goods{
		Name: "韩国面膜",
		Fact: true,
	}
	g2 := &Goods{
		Name: "雷碧",
		Fact: false,
	}

	s := Shopping(&KoreaShopping{})

	p := NewProxy(s)
	p.Buy(g1)
	p.Buy(g2)

}
