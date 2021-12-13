package kerusakan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	businessKerusakan "core-data/business/kerusakan"
)

type HTTPRequest struct{}

type KerusakanResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		ID             int    `json:"id"`
		JenisKerusakan string `json:"jenis_kerusakan"`
		LamaPengerjaan string `json:"lama_pengerjaan"`
		Harga          int    `json:"harga"`
		Version        int    `json:"version"`
	} `json:"data"`
}

type KerusakanDataResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             int    `json:"id"`
		JenisKerusakan string `json:"jenis_kerusakan"`
		LamaPengerjaan string `json:"lama_pengerjaan"`
		Harga          int    `json:"harga"`
		Version        int    `json:"version"`
	} `json:"data"`
}

func NewHTTPRequestRepository() *HTTPRequest {
	return &HTTPRequest{}
}

func (repo *HTTPRequest) GetListKerusakan() ([]businessKerusakan.Kerusakan, error) {
	var (
		reponse      KerusakanResponse
		kerusakan    businessKerusakan.Kerusakan
		allKerusakan []businessKerusakan.Kerusakan
	)

	url := "http://127.0.0.1:7070/v1/kerusakan"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &reponse)

	for _, data := range reponse.Data {
		kerusakan.ID = data.ID
		kerusakan.JenisKerusakan = data.JenisKerusakan
		kerusakan.LamaPengerjaan = data.LamaPengerjaan
		kerusakan.Harga = data.Harga
		kerusakan.Version = data.Version

		allKerusakan = append(allKerusakan, kerusakan)
	}

	return allKerusakan, nil
}

func (repo *HTTPRequest) GetKerusakan(id int) (*businessKerusakan.Kerusakan, error) {
	var (
		response  KerusakanDataResponse
		kerusakan businessKerusakan.Kerusakan
	)

	idString := strconv.Itoa(id)
	url := "http://127.0.0.1:7070/v1/kerusakan/id/" + idString

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &response)

	kerusakan.ID = response.Data.ID
	kerusakan.JenisKerusakan = response.Data.JenisKerusakan
	kerusakan.LamaPengerjaan = response.Data.LamaPengerjaan
	kerusakan.Harga = response.Data.Harga
	kerusakan.Version = response.Data.Version

	return &kerusakan, nil
}
