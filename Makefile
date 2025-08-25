# Password Generator Makefile

# 변수 정의
BINARY_NAME=password_generator
MAIN_FILE=main.go
BUILD_DIR=build

# 기본 타겟
.PHONY: all
all: build

# 빌드
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# 실행
.PHONY: run
run:
	@echo "Running $(BINARY_NAME)..."
	go run $(MAIN_FILE)

# 테스트
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# 정리
.PHONY: clean
clean:
	@echo "Cleaning build files..."
	@rm -rf $(BUILD_DIR)
	@rm -f password_data.json
	@echo "Clean completed"

# 설치 (Go 모듈 의존성 다운로드)
.PHONY: install
install:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Install completed"

# 린트 (golangci-lint가 설치되어 있는 경우)
.PHONY: lint
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install it first: https://golangci-lint.run/usage/install/"; \
	fi

# 포맷팅
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt $(MAIN_FILE)
	@echo "Format completed"

# 벤치마크 (테스트 파일이 있는 경우)
.PHONY: benchmark
benchmark:
	@echo "Running benchmarks..."
	go test -bench=. ./...

# 도움말
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build     - Build the binary"
	@echo "  run       - Run the program"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean build files"
	@echo "  install   - Install dependencies"
	@echo "  lint      - Run linter (requires golangci-lint)"
	@echo "  fmt       - Format code"
	@echo "  benchmark - Run benchmarks"
	@echo "  help      - Show this help message"
