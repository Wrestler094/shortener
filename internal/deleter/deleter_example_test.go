package deleter_test

import (
	"time"

	"github.com/Wrestler094/shortener/internal/deleter"
	"github.com/Wrestler094/shortener/internal/storage"
)

// Пример использования URLDeleter для асинхронного удаления URL
func Example() {
	// Создаем хранилище (в данном примере используем заглушку)
	var storage storage.IStorage

	// Создаем новый экземпляр URLDeleter с интервалом обновления 5 секунд
	d := deleter.NewURLDeleter(storage, 5*time.Second)

	// Запускаем фоновый процесс удаления
	d.StartBackgroundFlusher()

	// Добавляем URL в очередь на удаление
	d.QueueForDeletion("abc123", "user1")
	d.QueueForDeletion("def456", "user1")
	d.QueueForDeletion("ghi789", "user2")

	// URL будут автоматически удалены в фоновом режиме
	// через заданный интервал времени
}

// Пример демонстрирует, как использовать URLDeleter для удаления URL
// конкретного пользователя
func ExampleURLDeleter_QueueForDeletion() {
	var storage storage.IStorage
	d := deleter.NewURLDeleter(storage, 5*time.Second)
	d.StartBackgroundFlusher()

	// Добавляем несколько URL одного пользователя в очередь на удаление
	d.QueueForDeletion("abc123", "user1")
	d.QueueForDeletion("def456", "user1")
}
