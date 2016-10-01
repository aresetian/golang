
package controllers

import "github.com/astaxie/beego"
import "encoding/json"


type Login struct {
	beego.Controller
}

type LoginResult struct {
	User string
	Token string
}




// Use Get rather than Post so that we can simulate easier in the browser
func (this *Login) Get() {
   
 
   
        aSlice := LoginResult{}
        aSlice.User = "cgarcia"
        aSlice.Token = "4a65df454sd56f4a4sdf6a4sd654fa"
 
   s, _ := json.Marshal(aSlice) 
   this.Ctx.ResponseWriter.Write(s) 
   
}