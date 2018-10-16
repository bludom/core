#!/bin/bash

SERVER="localhost:8080/temperature"

example=curl -i -X POST 'http://gomano.de:8086/write?db=talk' --data-binary 'temperatur,host=arbeitslaptop value=32.0'


putTemperature() {
    local ID=$1
    local TEMP=$2
    curl -s -X PUT $SERVER -d "{\"device\" \"$(hostname)\": ,\"core\":$ID,\"temp\":$TEMP}"
}

checkTemperature() {
     local ID=$1
     local EXPECTEDTEMP=$2
     local SAVEDTEMP=$(curl -s -X GET "$SERVER/$ID")
     local SAVEDTEMP=${SAVEDTEMP:0:4}
     if [[ $EXPECTEDTEMP == $SAVEDTEMP ]]; then
        return 0
     fi
     return 1
}

sensors coretemp-isa-0000 | while read line; do
    if [[ "$line" =~ ^Core ]]; then
        ID=${line/:*/}
        ID=${ID#Core }

        TEMP=${line/*:/}
        TEMP=${TEMP/°*/}
        TEMP=${TEMP/+/}
        putTemperature $ID $TEMP
        if checkTemperature $ID $TEMP; then
            echo "${ID} saved successfully"
        else 
            echo "${ID} got wrong temp, expected ${TEMP}"
        fi
    fi
done



