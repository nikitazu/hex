@echo off
call .\build.cmd
if %errorlevel%==0 (
  pushd .\wrk
    ..\bin\hex.exe %*
  popd
)
