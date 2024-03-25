#!/usr/bin/bash

echo "Start"

if [ -d bin ]; then
    GOOS=windows GOARCH=386 go build -o ./bin/autoALCv3_windows_386.exe
    GOOS=windows GOARCH=amd64 go build -o ./bin/autoALCv3_windows_amd64.exe

    GOOS=linux GOARCH=386 go build -o ./bin/autoALCv3_linux_386
    GOOS=linux GOARCH=amd64 go build -o ./bin/autoALCv3_linux_amd64
    GOOS=linux GOARCH=arm go build -o ./bin/autoALCv3_linux_arm
fi

echo "Finish"
