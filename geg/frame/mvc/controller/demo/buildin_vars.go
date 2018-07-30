package demo

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/frame/gmvc"
)

type Order struct {
    gmvc.Controller
}

func init() {
    g.Server().BindController("/{.struct}/{.method}", &Order{})
}

func (o *Order) Index() {
    o.Response.Write("Order Index")
}

func (o *Order) List() {
    o.Response.Write("Order List")
}
