package controllers

import (
	"encoding/json"
	"test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type returnObj struct {
	successful bool
	id         int64
}

// Operations about Product
type ProductController struct {
	beego.Controller
}

// @Title CreateProduct
// @Description create product
// @Param	body		body 		models.Product	true		"body for product content"
// @Success 200 	returnObj
// @Failure 403 body is empty
// @router /AddProduct [Post]
func (p *ProductController) AddProduct() {
	log := logs.GetLogger()
	log.Println(p.Ctx.Input.RequestBody)
	var product models.Product
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &product)
	if err != nil {
		panic(err)
	}
	log.Println(product)
	var result returnObj
	// result.id, result.successful = models.AddProduct(product)
	p.Data["json"] = &result
	p.ServeJSON()
}
