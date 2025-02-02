#!/bin/bash

if [ "$EUID" -ne 0 ]; then 
    echo "You need root privileges for this program to work"
    exec sudo "$0" "$@"
    exit
fi

FRONTEND="./release/fiberOrca"
BACKEND="./release/network_analyzer"

$FRONTEND &
$BACKEND

