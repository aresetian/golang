
package controllers

import "github.com/astaxie/beego"
//import "encoding/json"


type Enterprise struct {
	beego.Controller
}

/*
 GeoLocalitation
*/
type EnterpriseResult struct {
	Sede []SedeResult
}

type SedeResult struct {
	Ciudad string
	Direccion string
	Telefonos []string
	Location string
}


// Use Get rather than Post so that we can simulate easier in the browser
func (this *Enterprise) Get() {
    var result EnterpriseResult
    
    phones1 := []string{"3256987", "(57)4578522", "1596357"}
    sede1 := SedeResult{"Cali", "Cll 14 # 64c 44", phones1,"1235454158474" }
    
    phones2 := []string{"63636987", "(44)4578522", "3566357"}
    sede2 := SedeResult{"Tulua","Carrera 65 # 12a 49", phones2,"5487552258885" }
    
    phones3 := []string{"6586987", "(54)4578522", "7777357"}
    sede3 := SedeResult{"Palmira","Avenida los cerros # 56 - 89", phones3,"8555854778956" }
    
    aSlice := []SedeResult{sede1, sede2, sede3}
    
    result.Sede = aSlice
	
	
	this.Data["json"] = result
	this.ServeJSON()
}