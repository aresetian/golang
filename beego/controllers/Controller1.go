package controllers

import "github.com/astaxie/beego"

type Controller1 struct {
	beego.Controller
}

type Controller1Result struct {
	Value1 string
	Value2 string
}

// Use Get rather than Post so that we can simulate easier in the browser
func (this *Controller1) Get() {
	var result Controller1Result
	valuel := "value 1"
	beego.Info(valuel)
	result.Value1 = valuel
	
	value2 := "value 2"
	beego.Info(value2)
	result.Value2 = value2
	
	this.Data["json"] = result
	this.ServeJSON()
}