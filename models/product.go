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
	Version   int64     `orm:"default(1)"`
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
// errorType 1 is updated DB by other transaction
func Purchase(purchaseOrder PurchaseOrder, transaction *orm.Ormer) (success bool, errorType int) {
	o := *transaction
	var product Product
	qs := o.QueryTable("product")
	err := qs.Filter("id", purchaseOrder.ProductID).Filter("quatity__gte", purchaseOrder.Quatity).One(&product)
	if err == orm.ErrMultiRows {
		// Have multiple records
		return false, 0
	}
	if err == orm.ErrNoRows {
		// No result
		return false, 0
	}
	product.Quatity = product.Quatity - purchaseOrder.Quatity
	num, err := qs.Filter("id", product.ProductID).Filter("Version__iexact", product.Version).Update(orm.Params{
		"Quatity": product.Quatity,
		"Version": product.Version + 1,
	})
	if num == 0 {
		return false, 1
	}
	return err == nil, 0
}

//Purchases array of order
func Purchases(purchaseOrderList []PurchaseOrder) bool {
	o := orm.NewOrm()
	optimisticLock := true
	for optimisticLock {
		err := o.Begin()
		if err != nil {
			return false
		}
		updatedDbFlag := false
		for _, purchaseOrder := range purchaseOrderList {
			result, errType := Purchase(purchaseOrder, &o)
			//Other transaction updated db. Need to rollback from start
			if errType == 1 {
				o.Rollback()
				updatedDbFlag = true
				break
			}
			if !result {
				o.Rollback()
				return false
			}
		}
		optimisticLock = updatedDbFlag
	}
	err := o.Commit()
	return err == nil
}
