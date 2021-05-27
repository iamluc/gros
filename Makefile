build:
	@go build -o gros

build-release:
	@mkdir -p dist
	@GOOS=linux GOARCH=amd64 go build -o ./dist/gros_linux_amd64
	@GOOS=linux GOARCH=arm64 go build -o ./dist/gros_linux_arm64
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/gros_linux_static_amd64
	@GOOS=darwin GOARCH=amd64 go build -o ./dist/gros_darwin_amd64
	@GOOS=darwin GOARCH=arm64 go build -o ./dist/gros_darwin_arm64
	@GOOS=freebsd GOARCH=amd64 go build -o ./dist/gros_freebsd_amd64
	@GOOS=windows GOARCH=amd64 go build -o ./dist/gros_windows_amd64.exe
