/* tick-tock.go */
package main

import (
        "fmt"
        "log"
        "net/http"
        //"time"
        //"os"
)

// Content for the main html page..
//src="http://ajax.googleapis.com/ajax/libs/jquery/1.3.2/jquery.min.js">
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
                        url: 'https://poc-golang-carlos-andres.c9.io/Controller2',
                        type: 'post',
                        crossDomain: true,
                        headers: {
                          'Access-Control-Allow-Origin': '*'
                        },
                        dataType: 'json',
                        data : "&processName=1",
                        success : function(data) {
                            alert(data);
                            alert(data.length)
                        }
                });
             </script>
           </body>
         </html>`

// handler for the main page.
func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, page)
}

func main() {
        http.HandleFunc("/time", handler)
        //https://docs.c9.io/docs/multiple-ports
        //log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
        fmt.Println("listening....")
        log.Fatal(http.ListenAndServe("0.0.0.0:8082", nil))
}   