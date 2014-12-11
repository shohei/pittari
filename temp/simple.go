// httpserver.go
package main

import (
	"flag"
	"net/http"
	"log"
	"time"
)

var port = flag.String("port", "8000", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")


func main(){
	sleep1_finished := make(chan bool)
    sleep2_finished := make(chan bool)

	go func () {
		flag.Parse()
		panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*root))))
		 sleep1_finished <- true //never reached
	}()

	go func(){
		// 2秒かかるコマンド
        log.Print("sleep2 started.")
        time.Sleep(5 * time.Second)
        log.Print("sleep2 finished.")
        sleep2_finished <- true
	}()

     <- sleep1_finished //never reached
     <- sleep2_finished

	 log.Print("finished") //never reached

}
