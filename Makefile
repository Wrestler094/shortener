# Имя бинарника (можно изменить под проект)
BINARY_NAME = shortener

# Путь к точке входа (main.go)
ENTRY = ./cmd/shortener/main.go

# Импорт-префикс вашего проекта для -local
LOCAL_MODULE = github.com/Wrestler094/shortener

.PHONY: all fmt imports build run

all: fmt imports build

# Форматирование кода с gofmt
fmt:
	@gofmt -s -w .

# Упорядочивание импортов, локальные импорты после внешних
imports:
	@goimports -w -local $(LOCAL_MODULE) .

# Сборка бинарного файла
build:
	@go build -o $(BINARY_NAME) $(ENTRY)

# Запуск приложения
run:
	@go run $(ENTRY)
