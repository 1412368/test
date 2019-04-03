package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Product obj
type Product struct {
	ProductID int64     `json:"productid" orm:"column(id);pk"`
	Quantity  int64     `json:"quantity"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Product))
	// p := Product{"1", "12"}
	// UserList["user_11111"] = &u
}

//AddProduct use to insert product into product table
//@Params product
func AddProduct(product Product) (id int64, successful bool) {
	o := orm.NewOrm()
	id, e := o.Insert(&product)
	return id, e == nil
}
