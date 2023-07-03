# Nama aplikasi
APP_NAME := api


# Build versi Windows
build-windows:
	@echo "Building Windows..."
	set GOOS=windows
	set GOARCH=amd64
	go build -o windows-$(APP_NAME).exe

# Build versi Linux
build-linux:
	@echo "Building Linux..."
	set GOOS=linux
	set GOARCH=amd64
	go build -o linux-$(APP_NAME)

# Build versi macOS
build-mac:
	@echo "Building macOS..."
	set GOOS=darwin
	set GOARCH=amd64
	go build -o mac-$(APP_NAME)

# Build untuk semua platform
build: build-windows build-linux build-mac


# Menjalankan aplikasi
run:
	go run main.go
