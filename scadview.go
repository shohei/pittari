package main

import (
	"flag"
	"net/http"
	"os/exec"
	"log"
	"os"
	"fmt"
)

var port = flag.String("port", "8000", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")
var scad_file = flag.String("f", "", "Define the input scad file")

func main(){
	flag.Parse()

	if(*scad_file==""){
		fmt.Println("#Usage: go run scadview.go -f <input_scad>")
		fmt.Println("[Error]: input file not specified.")
		os.Exit(0)
	}

	sleep1_finished := make(chan bool)
	sleep2_finished := make(chan bool)

	//Run static HTTP server
	go func () {
		log.Print("server lisening at localhost:8000")
		flag.Parse()
		panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*root))))
		sleep1_finished <- true //never reached
	}()

	//open viewer and watch change
	go func(){
		log.Print("opening stl viewer.")
		log.Print("scad view watching started.")
		sh:= "/bin/sh"
		opt := "-c"
		scad_view_bytes := make([]byte,0,40)
		scad_scripts := "./scripts/scadview.sh "
		scad_view_bytes = append(scad_view_bytes,scad_scripts...)
		scad_view_bytes = append(scad_view_bytes,*scad_file...)
		scadview := string(scad_view_bytes)

		cmd := exec.Command(sh,opt,scadview)
		stdout, err := cmd.Output()
		if err != nil {
			println(err.Error())
			return
		}
		print(string(stdout))
		sleep2_finished <- true
	}()

	<- sleep1_finished //never reached
	<- sleep2_finished //never reached

	log.Print("program finished") //never reached

}

