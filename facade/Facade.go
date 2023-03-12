package facade

import "fmt"

/**
  外观（Facade）模式
  主要是定义了一个高层接口,它包含了对各个子系统的引用，客户端可以通过它访问各个子系统的功能。

  如电商系统中商品详情页来说，商详页接口一般会涉及商品、库存、营销、商家等系统。
  为了避免让客户端调用接口拼凑出商详页的数据，一般商品组同学会提供商详页接口，通过该接口获取商详页的所有信息，返回给客户端。
*/

type GoodsDetail struct {
}

func (gd *GoodsDetail) GetGoodsDetailInfo() {
	fmt.Println("获取商品信息")
}

type GoodsStock struct {
}

func (gs *GoodsStock) GetGoodsStockInfo() {
	fmt.Println("获取商品库存信息")
}

type GoodsFactory struct {
}

func (gf *GoodsFactory) GetGoodsFactoryInfo() {
	fmt.Println("获取商品厂家信息")
}

type GoodsPromotion struct {
}

func (gp *GoodsPromotion) GetGoodsPromotionInfo() {
	fmt.Println("获取商品营销信息")
}

type GoodsFacade struct {
	GoodsDetail    *GoodsDetail
	GoodsStock     *GoodsStock
	GoodsFactory   *GoodsFactory
	GoodsPromotion *GoodsPromotion
}

func (f *GoodsFacade) GoodsInfo() {
	f.GoodsFactory.GetGoodsFactoryInfo()
	f.GoodsStock.GetGoodsStockInfo()
	f.GoodsPromotion.GetGoodsPromotionInfo()
	f.GoodsDetail.GetGoodsDetailInfo()
}
