package storage

type IStorage interface {
	// Init()
	Save(string, string)
	Get(string) (string, bool)
}
