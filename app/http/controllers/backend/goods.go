package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/app/services/backend"
)


func (ctrl *controller) Create(c *gin.Context) {
	var goodsForm backend.GoodsForm
	var goodsExtendForm backend.GoodsExtendForm
	var goodsSkuForm backend.GoodsSkuForm

	if err := c.ShouldBindJSON(&goodsForm); err != nil {
		ctrl.Fail(c, err)
		return
	}
	if err := c.ShouldBindJSON(&goodsExtendForm); err != nil {
		ctrl.Fail(c, err)
		return
	}
	if err := c.ShouldBindJSON(&goodsSkuForm); err != nil {
		ctrl.Fail(c, err)
		return
	}

	fmt.Println(goodsForm, goodsSkuForm, goodsExtendForm)
}
