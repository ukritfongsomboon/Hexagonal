package storage

type storageApp struct {
	disk string
}

func NewAppStorage(disk string) StorageApp {
	return storageApp{disk: disk}
}

func (sr storageApp) Write() error {
	return nil
}

func (sr storageApp) Read() error {
	return nil
}

func (sr storageApp) Delete() error {
	return nil
}
