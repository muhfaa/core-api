package kerusakan

type Kerusakan struct {
	ID             int    `json:"_" db:"id"`
	JenisKerusakan string `json:"jenis_kerusakan" db:"jenis_kerusakan"`
	LamaPengerjaan string `json:"lama_pengerjaan" db:"lama_pengerjaan"`
	Harga          int    `json:"harga" db:"harga"`
	Version        int    `json:"version" db:"version"`
}

func NewKerusakan(
	jenis_kerusakan string,
	lama_pengerjaan string,
	harga int,
) Kerusakan {

	return Kerusakan{
		JenisKerusakan: jenis_kerusakan,
		LamaPengerjaan: lama_pengerjaan,
		Harga:          harga,
		Version:        1,
	}
}
