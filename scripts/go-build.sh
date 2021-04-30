#!/bin/bash
# https://github.com/complexorganizations/go-build-script

# Pre-Checks system requirements
function installing-system-requirements() {
    if [ ! -x "$(command -v checksum)" ]; then
        echo "The application checksum was not found in the system. [ https://github.com/complexorganizations/checksum ]"
        exit
    elif [ ! -x "$(command -v go)" ]; then
        echo "The application go was not found in the system. [ https://go.dev ]"
        exit
    fi
}

# Run the function and check for requirements
installing-system-requirements

# Build for all the OS
function build-golang-app() {
    APPLICATION="code-formatter"
    VERSION="v1.0.0"
    if [ -n "$(ls ./*.go)" ]; then
        # Darwin
        GOOS=darwin GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-darwin-amd64 .
        GOOS=darwin GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-darwin-arm64 .
        # Linux
        GOOS=linux GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-linux-386 .
        GOOS=linux GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-linux-amd64 .
        GOOS=linux GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-linux-arm .
        GOOS=linux GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-linux-arm64 .
        GOOS=linux GOARCH=mips go build -o bin/${APPLICATION}-${VERSION}-linux-mips .
        GOOS=linux GOARCH=mips64 go build -o bin/${APPLICATION}-${VERSION}-linux-mips64 .
        GOOS=linux GOARCH=mips64le go build -o bin/${APPLICATION}-${VERSION}-linux-mips64le .
        GOOS=linux GOARCH=mipsle go build -o bin/${APPLICATION}-${VERSION}-linux-mipsle .
        GOOS=linux GOARCH=ppc64 go build -o bin/${APPLICATION}-${VERSION}-linux-ppc64 .
        GOOS=linux GOARCH=ppc64le go build -o bin/${APPLICATION}-${VERSION}-linux-ppc64le .
        GOOS=linux GOARCH=riscv64 go build -o bin/${APPLICATION}-${VERSION}-linux-riscv64 .
        GOOS=linux GOARCH=s390x go build -o bin/${APPLICATION}-${VERSION}-linux-s390x .
        # Windows
        GOOS=windows GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-windows-386.exe .
        GOOS=windows GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-windows-amd64.exe .
        GOOS=windows GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-windows-arm.exe .
        # Get SHA and put everything in a register.
        checksum "$PWD/bin"
    else
        echo "Error: The \".go\" files could not be found."
        exit
    fi
}

# Start the build
build-golang-app
