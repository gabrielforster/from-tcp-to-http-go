.PHONY: build test clean

TARGET_DIR = bin

build:
	@if [ ! -d "$(TARGET_DIR)" ]; then \
		mkdir -p "$(TARGET_DIR)"; \
	fi
	@echo "Building..."
	@go build -o $(TARGET_DIR)/main main.go

test:
	@echo "No tests yet"

clean:
	@echo "Cleaning up..."
	@rm -rf $(TARGET_DIR)
