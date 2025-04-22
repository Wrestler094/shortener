# Имя бинарника (можно изменить под проект)
BINARY_NAME = shortener

# Путь к точке входа (main.go)
ENTRY = ./cmd/shortener/main.go

# Импорт-префикс вашего проекта для -local
LOCAL_MODULE = github.com/Wrestler094/shortener

.PHONY: all fmt imports build run

all: fmt imports build

# Сборка бинарного файла
build:
	@go build -o $(BINARY_NAME) $(ENTRY)

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

# Очистка go модулей и зависимостей
tidy:
	@go mod tidy
