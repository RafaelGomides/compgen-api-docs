package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:DocumentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:PermissionsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServiceOrderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"] = append(beego.GlobalControllerRouter["compgen-api-docs/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
