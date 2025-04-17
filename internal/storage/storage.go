package storage

import "context"

type IStorage interface {
	Save(string, string) error
	Get(string) (string, bool)
	SaveBatch(map[string]string) error
	FindShortByOriginalURL(string) (string, error)
}

type IPingableStorage interface {
	Ping(ctx context.Context) error
}
