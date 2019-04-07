package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["test/controllers:ProductController"] = append(beego.GlobalControllerRouter["test/controllers:ProductController"],
        beego.ControllerComments{
            Method: "AddProduct",
            Router: `/AddProduct`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/controllers:ProductController"] = append(beego.GlobalControllerRouter["test/controllers:ProductController"],
        beego.ControllerComments{
            Method: "Purchases",
            Router: `/purchases`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
