package repos

// BaseRepository sss
type BaseRepository interface {
	Insert() (string, error)
	GetAll()
}
