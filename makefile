win:
	GOOS=windows GOARCH=amd64 go build --ldflags="-w -s" -o matchapi.exe
	upx matchapi.exe

build:
	GOOS=linux GOARCH=amd64 go build --ldflags="-w -s" -o matchapi
	upx matchapi