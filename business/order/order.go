package order

type State string

const (
	StateAntrian    State = "antrian"
	StateDikerjakan State = "dikerjakan"
	StateSelesai    State = "selesai"
)

type Order struct {
	ID             int    `json:"_" db:"id"`
	TeknisiID      int    `json:"teknisi_id" db:"teknisi_id"`
	KerusakanID    int    `json:"kerusakan_id" db:"kerusakan_id"`
	JenisHP        string `json:"jenis_hp" db:"jenis_hp"`
	JenisPlatform  string `json:"jenis_platform" db:"jenis_platform"`
	StatusService  State  `json:"status_service" db:"status_service"`
	LamaPengerjaan string `json:"lama_pengerjaan" db:"lama_pengerjaan"`
	Version        int    `json:"version" db:"version"`
	VersionTeknisi int    `json:"version_teknisi" db:"version_teknisi"`
}

type OrderSpecRequest struct {
	TeknisiID     int    `json:"teknisi_id"`
	KerusakanID   int    `json:"kerusakan_id"`
	JenisHP       string `json:"jenis_hp"`
	JenisPlatform string `json:"jenis_platform"`
	Version       int    `json:"version"`
}

type OrderSpec struct {
	TeknisiID      int    `json:"teknisi_id"`
	KerusakanID    int    `json:"kerusakan_id"`
	JenisHP        string `json:"jenis_hp"`
	JenisPlatform  string `json:"jenis_platform"`
	StatusService  State  `json:"status_service"`
	LamaPengerjaan string `json:"lama_pengerjaan"`
	Version        int    `json:"version"`
	VersionTeknisi int    `json:"version_teknisi"`
}

func NewOrder(
	teknisi_id int,
	kerusakan_id int,
	jenis_hp string,
	jenis_platform string,
	status_service State,
	lama_pengerjaan string,
	version int,
	version_teknisi int,
) OrderSpec {

	return OrderSpec{
		TeknisiID:      teknisi_id,
		KerusakanID:    kerusakan_id,
		JenisHP:        jenis_hp,
		JenisPlatform:  jenis_platform,
		StatusService:  status_service,
		LamaPengerjaan: lama_pengerjaan,
		Version:        1,
		VersionTeknisi: version_teknisi,
	}
}

type UpdateOrderStatusSpec struct {
	ID      int `json:"id"`
	Version int `json:"version"`
}

type UpdateOrderStatus struct {
	ID            int
	TeknisiID     int
	StatusService State
	Version       int
}

func NewUpdateStatus(
	id int,
	teknisi_id int,
	status_service State,
	version int,
) UpdateOrderStatus {

	if status_service == StateAntrian {
		status_service = StateDikerjakan
	} else if status_service == StateDikerjakan {
		status_service = StateSelesai
	}

	return UpdateOrderStatus{
		ID:            id,
		TeknisiID:     teknisi_id,
		StatusService: status_service,
		Version:       version,
	}
}
