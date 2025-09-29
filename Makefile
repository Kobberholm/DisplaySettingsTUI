# DisplaySettingsTUI Makefile

BINARY_NAME=display-settings-tui
BUILD_DIR=build
INSTALL_DIR=/usr/local/bin

# Try to find go in common locations
GO := $(shell which go 2>/dev/null || echo "/usr/local/go/bin/go")

.PHONY: build install install-user uninstall clean run

build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -buildvcs=true -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd
	@echo "Built: $(BUILD_DIR)/$(BINARY_NAME)"

install:
	@if [ -f $(BUILD_DIR)/$(BINARY_NAME) ]; then \
		install -Dm755 $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME); \
		echo "Installed: $(INSTALL_DIR)/$(BINARY_NAME)"; \
	else \
		echo "Run 'make build' first, then 'sudo make install'"; \
		exit 1; \
	fi

install-user:
	@if [ -f $(BUILD_DIR)/$(BINARY_NAME) ]; then \
		mkdir -p ~/.local/bin; \
		install -Dm755 $(BUILD_DIR)/$(BINARY_NAME) ~/.local/bin/$(BINARY_NAME); \
		echo "Installed: ~/.local/bin/$(BINARY_NAME)"; \
		echo "Make sure ~/.local/bin is in your PATH"; \
	else \
		echo "Run 'make build' first, then 'make install-user'"; \
		exit 1; \
	fi

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstalled"

clean:
	rm -rf $(BUILD_DIR)
	@echo "Cleaned"

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)