package backend

type GoodsForm struct {
	Id uint `form:"id" json:"id"`
	Title string `form:"title" json:"title" binding:"required"`
	Desc string `form:"desc" json:"desc" binding:"required"`
	SellPrice float32 `json:"sell_price"`
	OriginPrice float32 `json:"origin_price"`
	IsSell uint8 `json:"is_sell" binding:"required"`
}

type GoodsExtendForm struct {
	GoodsId uint `form:"id" json:"goods_id"`
	Images []string `form:"images" json:"images" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Spec string `form:"spec" json:"spec" binding:"required"`
}

type GoodsSkuForm struct {
	GoodsId uint `form:"id" json:"goods_id" binding:"required"`
	SpecData string `form:"spec_data" json:"spec_data" binding:"required"`
	SellPrice float32
	OriginPrice float32
	Sku string
	Image string
}



