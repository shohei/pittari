#!/bin/zsh
if [ $# -ne 1 ] ; then
  echo "#Usage: write.sh <scad file>"
  exit -1
fi
fswatch $@ | xargs -n1 -I{} ./exec_jscad.sh {} 
