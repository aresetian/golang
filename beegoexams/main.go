package main
    
import "github.com/astaxie/beego"
import "golang/beego/controllers"

func main(){
    beego.Router("/", &controllers.Controller1{})
    beego.Router("/Controller2", &controllers.Controller2{})
    beego.Router("/Controller3", &controllers.Controller3{})
    beego.Router("/Controller4", &controllers.Controller4{})
    beego.Run()
}