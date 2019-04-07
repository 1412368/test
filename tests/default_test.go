package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"test/db"
	_ "test/routers"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func SeedDB() {
	db.InitDatabase()
	var jsonStr = []byte(`
		{
		  "ProductID": 1,
		  "Quatity": 10
		}`)
	var jsonStr2 = []byte(`
		{
		  "ProductID": 2,
		  "Quatity": 5
		}`)
	r, _ := http.NewRequest("POST", "/v1/product/AddProduct", bytes.NewBuffer(jsonStr))
	r2, _ := http.NewRequest("POST", "/v1/product/AddProduct", bytes.NewBuffer(jsonStr2))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.BeeApp.Handlers.ServeHTTP(w, r2)
}
func TestPostPurchasesCase1(t *testing.T) {
	SeedDB()
	var jsonStr = []byte(`[
		{
		  "ProductID": 1,
		  "Quatity": 2
		},{
		  "ProductID": 2,
		  "Quatity": 1
		}
	      ]`)
	r, _ := http.NewRequest("POST", "/v1/product/purchases", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestPostPurchasesCase2(t *testing.T) {
	SeedDB()
	var jsonStr = []byte(`[
		{
		  "ProductID": 1,
		  "Quatity": 2
		},{
		  "ProductID": 2,
		  "Quatity": 6
		}
	      ]`)
	r, _ := http.NewRequest("POST", "/v1/product/purchases", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 422", func() {
			So(w.Code, ShouldEqual, 422)
		})
	})
}
func CreatRequest(jsonStr []byte, requests chan httptest.ResponseRecorder) {
	r, _ := http.NewRequest("POST", "/v1/product/purchases", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	requests <- *w
}
func TestPostPurchasesCase3(t *testing.T) {
	SeedDB()
	requests := make(chan httptest.ResponseRecorder, 2)
	var jsonStr = []byte(`[
		{
		  "ProductID": 1,
		  "Quatity": 2
		},{
		  "ProductID": 2,
		  "Quatity": 1
		}
	      ]`)

	var jsonStr2 = []byte(`[
		{
		  "ProductID": 1,
		  "Quatity": 2
		},{
		  "ProductID": 2,
		  "Quatity": 1
		}
	      ]`)
	go CreatRequest(jsonStr, requests)
	go CreatRequest(jsonStr2, requests)
	w1 := <-requests
	w2 := <-requests
	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w1.Code, ShouldEqual, 200)
			So(w2.Code, ShouldEqual, 200)
		})
	})
}
