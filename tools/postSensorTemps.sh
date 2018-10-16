#!/bin/bash

SERVER="localhost:8080/temperature"
DEVICE=${1:-"arm"}

calc() {
    awk "BEGIN { print $* }"
}

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


if [ "${DEVICE,,}" == "arm" ]; then
    TEMP=$(cat /sys/class/thermal/thermal_zone0/temp)
    TEMP=$(calc ${TEMP}/1000)
    ID=0
    putTemperature "$ID" "$TEMP"
else
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
fi
