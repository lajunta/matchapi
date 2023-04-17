win:
	GOOS=windows GOARCH=amd64 go build --ldflags="-w -s" -o bin/matchapi.exe
	upx bin/matchapi.exe

build:
	GOOS=linux GOARCH=amd64 go build --ldflags="-w -s" -o bin/matchapi
	upx bin/matchapi
