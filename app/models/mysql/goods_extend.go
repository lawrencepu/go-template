package mysql

type GoodsExtend struct {
	Model
	GoodsId uint `gorm:"comment:'商品id'" json:"goods_id"`
	Images string `gorm:"comment:'相册-json数组'" json:"images"`
	Content string `json:"content"`
	Spec string `json:"spec"`
}