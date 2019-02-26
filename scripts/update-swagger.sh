#!/bin/bash
set -e

ApiRepo="$HOME/source/azure-rest-api-specs"

serviceFolders=$(ls -d $ApiRepo/specification/*/)

for serviceFolder in $serviceFolders
do
    echo $serviceFolder
    { 
        swaggerFolders=$(ls -d ${serviceFolder}resource-manager/*/)
    } || {
        swaggerFolders=""
    }
    for swaggerFolder in $swaggerFolders
    do
        echo "    $swaggerFolder"
        ls -d ${swaggerFolder}stable/*/
    done
    echo ""
done
