build:
	@go build -o gros

build-windows:
	@GOOS=windows GOARCH=amd64 go build -o gros_windows_amd64.exe
