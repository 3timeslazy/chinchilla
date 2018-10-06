package storage

// Storage describes methods
// necessary for storing urls
type Storage interface {
	Keep(short, longURL string) error
	GetLongByShort(short string) (string, error)
	GetShortByLong(longURL string) (string, error)
}
