package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductList contains array of product
type ProductList struct {
	products map[string]*Product
}

//PurchaseOrder use to parse input
type PurchaseOrder struct {
	ProductID int64
	quality   int64
}

//Product obj
type Product struct {
	ProductID int64     `json:"ProductID" orm:"column(id);pk"`
	Quality   int64     `json:"Quality"`
	Created   time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `json:"-" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Product))
	// p := Product{"1", "12"}
	// UserList["user_11111"] = &u
}

//FindProduct find one product in db
//@Params productID
func FindProduct(productID int64) (product Product, successful bool) {
	o := orm.NewOrm()
	result := Product{ProductID: productID}
	err := o.Read(&product)
	return result, err == nil
}

//AddProduct use to insert product into product table
//@Params product
func AddProduct(product Product) (id int64, successful bool) {
	o := orm.NewOrm()
	id, e := o.Insert(&product)
	return id, e == nil
}

//Purchase product
func Purchase(purchaseOrder PurchaseOrder) bool {
	o := orm.NewOrm()
	product := Product{ProductID: purchaseOrder.ProductID}
	if o.Read(&product) != nil {
		return false
	}
	if product.Quality < purchaseOrder.quality {
		return false
	}
	product.Quality = product.Quality - purchaseOrder.quality
	_, err := o.Update(&product)
	return err == nil
}
