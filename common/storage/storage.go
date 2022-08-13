package storage

type AppStorage interface {
	Write(string, []byte) error
	Read(string) error
	Delete(string) error
}
