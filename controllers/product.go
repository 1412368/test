package controllers

import (
	"encoding/json"
	"test/models"

	"github.com/astaxie/beego"
)

//ProductController  Operations about Product
type ProductController struct {
	beego.Controller
}

//AddProduct product
// @Title AddProduct
// @Description create product
// @Param	body		body 		models.Product	true		"body for product content"
// @Success 200 	{object} models.Status
// @Failure 403 body is empty
// @router /AddProduct [Post]
func (p *ProductController) AddProduct() {
	var product models.Product
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &product)
	if err != nil {
		panic(err)
	}
	_, result := models.AddProduct(product)
	returnObj := models.Status{Successful: result}
	p.Data["json"] = &returnObj
	p.ServeJSON()
}

//Purchases array of product
// @Title Purchases product
// @Description buy product
// @Param	body		body 		[]models.PurchaseOrder	true		"body for purchase"
// @Success 200 	{object} models.Status
// @Failure 403 	body not found
// @Failure 422		not enough stock
// @router /purchases [Post]
func (p *ProductController) Purchases() {
	var purchaseOrderList []models.PurchaseOrder
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &purchaseOrderList)
	var returnObj models.Status
	if err != nil {
		p.Ctx.Output.SetStatus(403)
		returnObj = models.Status{Successful: false}
	} else {
		result := models.Purchases(purchaseOrderList)
		returnObj = models.Status{Successful: result}
		if !result {
			p.Ctx.Output.SetStatus(422)
		}
	}
	p.Data["json"] = &returnObj
	p.ServeJSON()
}
