package storage

type StorageApp interface {
	Write() error
	Read() error
	Delete() error
}
