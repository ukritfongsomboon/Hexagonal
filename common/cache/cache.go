package cache

type AppCache interface {
	Get(string) (*string, error)
	Set(string, string) error
	Clear() error
}
