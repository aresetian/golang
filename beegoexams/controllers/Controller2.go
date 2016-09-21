
package controllers

import "github.com/astaxie/beego"
import "encoding/json"


type Controller2 struct {
	beego.Controller
}

type Controller2Result struct {
	Accommodation []string
	Vehicle []string
}

// Use Get rather than Post so that we can simulate easier in the browser
func (this *Controller2) Get() {
    var result Controller2Result
    
    aSlice := []string{"House", "Apartment", "Hostel"}
    bSlice := []string{"Car", "Moto", "Airplane"}
    
    result.Accommodation = aSlice
	result.Vehicle = bSlice
	
	this.Data["json"] = result
	this.ServeJSON()
}


func (this *Controller2) Post() {
   var result Controller2Result
   
   aSlice := []string{"House", "Apartment", "Hostel"}
   bSlice := []string{"Car", "Moto", "Airplane"}
   result.Accommodation = aSlice
   result.Vehicle = bSlice

   s, _ := json.Marshal(result) 
   this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
   this.Ctx.ResponseWriter.Write(s) 
   
}