#!/bin/zsh
echo $@ | sed -e 's?'`pwd`/'??' -e 's/\.scad//g'| xargs -n 1 -I{} openjscad {}.scad -o {}.stl 2>/dev/null 1>/dev/null
echo $?
