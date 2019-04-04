package controllers

import (
	"encoding/json"
	"test/models"

	"github.com/astaxie/beego"
)

type returnObj struct {
	successful bool
	id         int64
}

//ProductController  Operations about Product
type ProductController struct {
	beego.Controller
}

//AddProduct product
// @Title AddProduct
// @Description create product
// @Param	body		body 		models.Product	true		"body for product content"
// @Success 200 	returnObj
// @Failure 403 body is empty
// @router /AddProduct [Post]
func (p *ProductController) AddProduct() {
	var product models.Product
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &product)
	if err != nil {
		panic(err)
	}
	var result returnObj
	result.id, result.successful = models.AddProduct(product)
	p.Data["json"] = &result
	p.ServeJSON()
}

//Purchase product
// @Title Purchase product
// @Description buy product
// @Param	body		body 		models.PurchaseOrder	true		"body for purchase"
// @Success 200 	bool
// @Failure 403 body is empty
// @router /purchase [Post]
func (p *ProductController) Purchase() {
	var purchaseOrder models.PurchaseOrder
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &purchaseOrder)
	if err != nil {
		panic(err)
	}
	result := models.Purchase(purchaseOrder)
	p.Data["json"] = &result
	p.ServeJSON()
}
