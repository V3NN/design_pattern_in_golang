package facade

import "testing"

func TestFacade(t *testing.T) {
	g := new(GoodsFacade)
	g.GoodsInfo()
}
