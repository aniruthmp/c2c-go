#!/usr/bin/env bash

for i in {1..10}
do
    curl "https://c2c-demo.apps.sbx.pcf.evpsradiology.com/c2c"

    printf "\n"
    for load in $(seq 1 15); do
    echo -ne "Slept $load seconds ..\r"
    sleep 1s
    done

done
