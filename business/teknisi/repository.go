package teknisi

type Repository interface {
	GetListTeknisi() ([]Teknisi, error)
	AddAntrian(UpdateJumlahAntrian) (bool, error)
	EraseAntrian(UpdateJumlahAntrian) (bool, error)
}
