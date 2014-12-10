if [ $# -ne 1 ]; then
  echo "#Usage: scadview.sh <SCAD FILE NAME>"
  exit -1
fi
EXT=$(echo $@ | sed -e 's/.*\.//g')
if [ "$EXT" != "scad" ]; then
   echo "Input file is not a scad file"
   exit -1
fi
FILE=$(echo $@ | sed -e 's/\.scad//g')
sed -i -e 's?loadSTL(".*\.stl")?loadSTL("/'$FILE'.stl")?g' index.html
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --app="http://localhost:8000"
./watch.sh $@
