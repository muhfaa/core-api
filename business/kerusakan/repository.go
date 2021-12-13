package kerusakan

type Repository interface {
	GetListKerusakan() ([]Kerusakan, error)
	GetKerusakan(id int) (*Kerusakan, error)
}
