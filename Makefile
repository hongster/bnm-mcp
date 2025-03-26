APP_EXECUTABLE=bnm-mcp

.PHONY: build

build: cmd/stdio/main.go
	# Build Linux executable
	@echo "Building Linux executable..."
	GOOS=linux GOARCH=amd64 go build -o build/linux_amd64/${APP_EXECUTABLE} cmd/stdio/main.go

	# Build MacOS executable
	@echo "Building MacOS executable..."
	GOOS=darwin GOARCH=amd64 go build -o build/darwin_amd64/${APP_EXECUTABLE} cmd/stdio/main.go
	
	# Build Windows executable
	@echo "Buiilding Windows executable..."
	GOOS=windows GOARCH=amd64 go build -o build/windows_amd64/${APP_EXECUTABLE}.exe cmd/stdio/main.go

run:
	@echo "Running the MCP STDIO server..."
	go run cmd/stdio/main.go

clean:
	@echo "Cleaning up build artifacts..."
	go clean
	rm -rf build/*

test:
	go test ./...