package main

import "os/exec"

func main() {
    sh:= "/bin/sh"
    scadview := "stlview.sh"
    scad := "test.scad"

    cmd := exec.Command(sh,scadview,scad)
    stdout, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }

    print(string(stdout))
}

