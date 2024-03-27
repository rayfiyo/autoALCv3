#!/usr/bin/bash

echo "Start"

if [ -d bin ]; then
    GOOS=windows GOARCH=386 go build -o "./bin/autoALCv3_windows_386.exe"
    GOOS=windows GOARCH=amd64 go build -o "./bin/autoALCv3_windows_amd64.exe"
    GOOS=windows GOARCH=arm go build -o "./bin/autoALCv3_windows_arm.exe"
    GOOS=windows GOARCH=arm64 go build -o "./bin/autoALCv3_windows_arm64.exe"

    GOOS=linux GOARCH=386 go build -o "./bin/autoALCv3_linux_386"
    GOOS=linux GOARCH=amd64 go build -o "./bin/autoALCv3_linux_amd64"
    GOOS=linux GOARCH=arm go build -o "./bin/autoALCv3_linux_arm"
    GOOS=linux GOARCH=arm64 go build -o "./bin/autoALCv3_linux_arm64"

    GOOS=darwin GOARCH=amd64 go build -o "./bin/autoALCv3_mac_amd64"
    GOOS=darwin GOARCH=arm64 go build -o "./bin/autoALCv3_mac_arm64"
fi

echo "Finish"
