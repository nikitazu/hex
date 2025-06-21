@echo off
call .\build.cmd
if %errorlevel%==0 (
  pushd .\wrk
    ..\bin\hex.exe a.txt b.txt a.png b.png %*
  popd
)
