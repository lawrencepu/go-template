package mysql

type GoodsSku struct {
	Model
	GoodsId uint `gorm:"comment:'商品id'" json:"goods_id"`
	SpecData string `gorm:"comment:'规格数据，json'" json:"spec_data"`
	SellPrice	float32 `gorm:"comment:'售价'" json:"sell_price"`
	OriginPrice float32 `gorm:"comment:'原价'" json:"origin_price"`
	Sku string `gorm:"comment:'sku'" json:"sku"`
}