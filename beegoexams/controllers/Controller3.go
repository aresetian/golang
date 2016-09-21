
package controllers

import "github.com/astaxie/beego"
import "encoding/json"


type Controller3 struct {
	beego.Controller
}


// Use Get rather than Post so that we can simulate easier in the browser
func (this *Controller3) Get() {
 
    
   aSlice := []string{"foo", "bar"}
   bSlice := []string{"baz", "qux"}

   slice := []interface{}{aSlice, bSlice}
 
   s, _ := json.Marshal(slice) 
   this.Ctx.ResponseWriter.Write(s) 
}

func (this *Controller3) Post() {
 
    
   aSlice := []string{"foo", "bar"}
   bSlice := []string{"baz", "qux"}

   slice := []interface{}{aSlice, bSlice}
 
   s, _ := json.Marshal(slice) 
   this.Ctx.ResponseWriter.Write(s) 
}