package main
    
import "github.com/astaxie/beego"
import "golang/beegoexams/controllers"
import "github.com/astaxie/beego/plugins/cors"

/**
To execute the code, firts download 

*/

func main(){
    beego.Router("/News", &controllers.News{})
    beego.Router("/Login", &controllers.Login{})
    beego.Router("/Enterprise", &controllers.Enterprise{})
    beego.Router("/Exams", &controllers.Exams{})
    
    /** 
    https://stackoverflow.com/questions/35817926/cors-fetch-from-firefox-to-beego-server-stops-after-pre-flight
    */
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
  AllowOrigins:     []string{"*"},
  AllowMethods:     []string{"*"},
  AllowHeaders:     []string{"Origin"},
  ExposeHeaders:    []string{"Content-Length"},
  //Access-Control-Allow-Headers: []string{"Origin", "Content-Type"},
  AllowCredentials: true,
  
  
}))

    beego.Run()
}