@echo off
pushd .\src
  call go build -o ..\bin\hex.exe .
popd
