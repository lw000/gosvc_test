set prjPath=%cd%
echo %prjPath%
cd ../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=windows
cd %prjPath%
go build -o gosvc.exe -v -ldflags="-s -w"