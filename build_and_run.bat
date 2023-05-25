@echo off
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o launcher.exe
launcher.exe --verbose
