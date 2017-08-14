SET GOROOT=%cd%\golang

SET GOPATH=%cd%\golang\external;%cd%\..\server

SET PATH=%PATH%;%cd%\golang\bin;%cd%\gcc64\bin

cd ..\server\src\logic

go build

logic.exe

pause