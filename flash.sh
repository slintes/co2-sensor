#!/bin/zsh

sudo GOPATH=$GOPATH:$(pwd) tinygo flash -target=arduino-nano33 main.go

echo "Connect to USB port for debugging with 'sudo screen /dev/ttyACM0 9600'"