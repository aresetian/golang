
package controllers

import "github.com/astaxie/beego"


var page =
        `<html>
           <head>
             <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script>
             <style> 
               div {
                 font-family: "Times New Roman", Georgia, Serif;
                 font-size: 2em;
                 width: 13.3em;
                 padding: 8px 8px; 
                 border: 2px solid #2B1B17;
                 border-radius: 10px;
                 color: #2B1B17;
                 text-shadow: 1px 1px #E5E4E2;
                 background: #FFFFFF;
               }
             </style>
           </head>
           <body>
             <h2>Go Timer (ticks every second!), the system clock</h2>
             <div id="output"></div>
             jojo
             <script type="text/javascript">
             $.ajax({
                        url: '/Controller2',
                        type: 'post',
                        dataType: 'json',
                        success : function(data) {
                            alert(JSON.stringify(data));
                            alert(data);
                            alert(data[0]);    // ["foo","bar"]
                            alert(data[1]);    // ["baz","qux"]
                            alert(data.length)
                        }
                });
             </script>
           </body>
         </html>`
         
type Controller4 struct {
	beego.Controller
}


// Use Get rather than Post so that we can simulate easier in the browser
func (this *Controller4) Get() {
 
    
   this.Ctx.ResponseWriter.Write([]byte(page)) 
}
