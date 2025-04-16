package storage

import "context"

type IStorage interface {
	Save(string, string)
	Get(string) (string, bool)
}

type IPingableStorage interface {
	Ping(ctx context.Context) error
}
