package teknisi

type Teknisi struct {
	ID            int    `json:"_" db:"id"`
	FullName      string `json:"full_name" db:"full_name"`
	Specialist    string `json:"specialist" db:"specialist"`
	Platform      string `json:"platform" db:"platform"`
	JumlahAntrian int    `json:"jumlah_antrian" db:"jumlah_antrian"`
	Version       int    `json:"version" db:"version"`
}

func NewTeknisi(
	id int,
	full_name string,
	specialist string,
	platform string,
	jumlah_antrian int,
	version int,
) Teknisi {

	return Teknisi{
		ID:            id,
		FullName:      full_name,
		Specialist:    specialist,
		Platform:      platform,
		JumlahAntrian: jumlah_antrian,
		Version:       1,
	}
}

type TeknisiUpdateSpec struct {
	ID         int
	FullName   string
	Specialist string
	Platform   string
	Version    int
}

func NewTeknisiUpdate(
	id int,
	full_name string,
	specialist string,
	platform string,
	version int,
) TeknisiUpdateSpec {

	return TeknisiUpdateSpec{
		ID:         id,
		FullName:   full_name,
		Specialist: specialist,
		Platform:   platform,
		Version:    version,
	}
}

// Update State in Loan
type UpdateJumlahAntrian struct {
	ID      int
	Version int
}

func NewUpdateJumlahAntrian(
	id int,
	version int,
) UpdateJumlahAntrian {

	return UpdateJumlahAntrian{
		ID:      id,
		Version: version,
	}
}
