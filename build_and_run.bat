@echo off
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
echo Building...
go build -o .server/launcher.exe
cd .server
echo Running...
launcher.exe
@REM launcher.exe --verbose
