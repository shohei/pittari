package main

import "github.com/codeskyblue/go-sh"

func main() {
	msg,err := sh.Command("../scadview.sh","../hoge.scad").Output()

	if err != nil {
      println(err.Error())
	  return
	}
	print(string(msg))
}

