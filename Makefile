# Имя бинарника (можно изменить под проект)
BINARY_NAME = shortener

# Путь к точке входа (main.go)
ENTRY = ./cmd/shortener/main.go

# Импорт-префикс вашего проекта для -local
LOCAL_MODULE = github.com/Wrestler094/shortener

# Директория с исходными .proto файлами
PROTO_SRC_DIR := internal/grpc/proto
# Директория, куда будут сгенерированы Go файлы из .proto файлов
PROTO_OUT_DIR := internal/grpc/pb

# Параметры сборки
VERSION ?= $(shell git describe --tags --always --dirty)
DATE    ?= $(shell date +%Y-%m-%dT%H:%M:%S)
COMMIT  ?= $(shell git rev-parse --short HEAD)
LDFLAGS = -X 'main.buildVersion=$(VERSION)' -X 'main.buildDate=$(DATE)' -X 'main.buildCommit=$(COMMIT)'

.PHONY: all fmt imports build run

all: fmt imports build

# Сборка бинарного файла
build:
	@go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) $(ENTRY)

rebuild: clean build

# Запуск приложения
run:
	@go run $(ENTRY)

# Удаление скомпилированного бинарника
clean:
	@rm -f $(BINARY_NAME)

# Форматирование кода с gofmt
fmt:
	@gofmt -s -w .

# Упорядочивание импортов, локальные импорты после внешних
imports:
	@goimports -w -local $(LOCAL_MODULE) .

# Запуск всех тестов
test:
	@go test -v ./...

# Генерация Go-кода из proto-файлов
# Требует установленного protoc и go-grpc плагина
proto:
	@mkdir -p $(PROTO_OUT_DIR)
	@protoc \
		--go_out=$(PROTO_OUT_DIR) \
		--go-grpc_out=$(PROTO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=$(PROTO_SRC_DIR) \
		$(PROTO_SRC_DIR)/*.proto

# Очистка go модулей и зависимостей
tidy:
	@go mod tidy

version:
	@echo "Version: $(VERSION)"
	@echo "Date:    $(DATE)"
	@echo "Commit:  $(COMMIT)"
