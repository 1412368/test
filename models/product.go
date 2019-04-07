package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductList contains array of product
type ProductList struct {
	products map[string]*Product
}

//PurchaseOrderList contains array of PurchaseOrder
type PurchaseOrderList struct {
	products []PurchaseOrder
}

//PurchaseOrder use to parse input
type PurchaseOrder struct {
	ProductID int64
	Quatity   int64
}

//Product obj
type Product struct {
	ProductID int64     `json:"ProductID" orm:"column(id);pk"`
	Quatity   int64     `json:"Quatity"`
	Created   time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `json:"-" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Product))
	// seedDB()
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
func Purchase(purchaseOrder PurchaseOrder, transaction *orm.Ormer) bool {
	o := *transaction
	var product Product
	qs := o.QueryTable("product")
	err := qs.Filter("id", purchaseOrder.ProductID).Filter("quatity__gte", purchaseOrder.Quatity).One(&product)
	if err == orm.ErrMultiRows {
		// Have multiple records
		return false
	}
	if err == orm.ErrNoRows {
		// No result
		return false
	}
	product.Quatity = product.Quatity - purchaseOrder.Quatity
	_, err = o.Update(&product)
	return err == nil
}

//Purchases array of order
func Purchases(purchaseOrderList []PurchaseOrder) bool {
	o := orm.NewOrm()
	_, err := o.Raw("START TRANSACTION").Exec()
	if err != nil {
		return false
	}
	for _, purchaseOrder := range purchaseOrderList {
		result := Purchase(purchaseOrder, &o)
		if !result {
			o.Raw("ROLLBACK").Exec()
			return false
		}
	}
	o.Raw("COMMIT").Exec()
	return true
}
