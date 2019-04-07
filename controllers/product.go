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

//Purchases array of product
// @Title Purchases product
// @Description buy product
// @Param	body		body 		[]models.PurchaseOrder	true		"body for purchase"
// @Success 200 	bool
// @Failure 403 body is empty
// @router /purchases [Post]
func (p *ProductController) Purchases() {
	var purchaseOrderList []models.PurchaseOrder
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &purchaseOrderList)
	if err != nil {
		panic(err)
	}
	result := models.Purchases(purchaseOrderList)
	p.Data["json"] = &result
	p.ServeJSON()
}
