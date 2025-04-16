package storage

var (
	Storage IStorage
)

type IStorage interface {
	// Init()
	Save(string, string)
	Get(string) (string, bool)
}
