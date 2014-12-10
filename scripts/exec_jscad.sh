#!/bin/zsh
STATUS=$(./scripts/test_scad.sh $@)
if [ $STATUS -ne 0 ]; then
  echo "test failed"
  exit 0
fi
echo $@ | sed -e 's?'`pwd`/'??' -e 's/\.scad//g'|xargs -n 1 -I{} openjscad {}.scad -o {}.stl && ./scripts/reload_stl.sh || echo "converting failed"
