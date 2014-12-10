chrome-cli list tabs| grep "STL viewer" | sed -e 's/:.*//g' -e 's/\[//g'| chrome-cli reload
