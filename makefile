.PHONY: build install clean

# Build the application
build:
	@echo "Building pump..."
	go build -o pump

# Install the built binary
install: build
	@echo "Installing pump..."
	cp pump /usr/local/bin/ || (mkdir -p ~/.local/bin && cp pump ~/.local/bin/)
	@echo "Pump installed successfully"

# Clean up binaries
clean:
	@echo "Cleaning up..."
	rm -f pump
