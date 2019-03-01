package repos

// BaseRepository sss
type BaseRepository interface {
	FindByID() ([]byte, error)
}
