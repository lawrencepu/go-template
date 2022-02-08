package mysql

type Goods struct {
	Model
	Title string `gorm:"comment:'商品标题'" json:"user_agent"`
	Desc string `gorm:"comment:'商品描述'" json:"desc"`
	SellPrice	float32 `gorm:"comment:'售价'" json:"sell_price"`
	OriginPrice float32 `gorm:"comment:'原价'" json:"origin_price"`
	IsSell uint8 `gorm:"comment:'是否上架;1-上架;2-下架'" json:"is_sell"`
}
