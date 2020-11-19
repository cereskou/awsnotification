@echo off

set GOOS=linux
set GOARCH=amd64

echo Build ...
go build -ldflags="-s -w" -o main
if %errorlevel% neq 0 (
    echo Failed.
    goto :EOF
)

echo Make zip ...
build-lambda-zip.exe --output function.zip main
echo Done.
