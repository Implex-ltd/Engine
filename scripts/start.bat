@echo off
set "directory=C:\Users\arm\Desktop\MYBROWSER\gologin\prof"

cd ../cmd/engine

:loop
rd /s /q "%directory%"
mkdir "%directory%"

go run . engine

timeout /t 1 /nobreak > NUL
goto loop
