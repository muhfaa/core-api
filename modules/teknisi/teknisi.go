package teknisi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	businessTeknisi "core-data/business/teknisi"

	"github.com/labstack/gommon/log"
)

type HTTPRequest struct{}

type TeknisiResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		ID            int    `json:"id"`
		FullName      string `json:"full_name"`
		Specialist    string `json:"specialist"`
		Platform      string `json:"platform"`
		JumlahAntrian int    `json:"jumlah_antrian"`
		Version       int    `json:"version"`
	} `json:"data"`
}

type UpdateAntrianResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID int `json:"id"`
	} `json:"data"`
}

func NewHTTPRequestRepository() *HTTPRequest {
	return &HTTPRequest{}
}

// GetListTeknisi
func (repo *HTTPRequest) GetListTeknisi() ([]businessTeknisi.Teknisi, error) {
	var (
		reponse    TeknisiResponse
		teknisi    businessTeknisi.Teknisi
		allTeknisi []businessTeknisi.Teknisi
	)

	url := "http://127.0.0.1:7070/v1/teknisi"
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
		teknisi.ID = data.ID
		teknisi.FullName = data.FullName
		teknisi.Specialist = data.Specialist
		teknisi.Platform = data.Platform
		teknisi.JumlahAntrian = data.JumlahAntrian
		teknisi.Version = data.Version

		allTeknisi = append(allTeknisi, teknisi)
	}

	return allTeknisi, nil
}

// AddAntrian
func (repo *HTTPRequest) AddAntrian(updateSpec businessTeknisi.UpdateJumlahAntrian) (bool, error) {
	var response UpdateAntrianResponse

	client := &http.Client{
		Timeout: time.Second * time.Duration(100),
	}

	bodyPayload := new(bytes.Buffer)
	if err := json.NewEncoder(bodyPayload).Encode(updateSpec); err != nil {
		log.Error(err)
		return false, nil
	}

	url := "http://127.0.0.1:7070/v1/teknisi/add/antrian"
	request, err := http.NewRequest(http.MethodPut, url, bodyPayload)
	if err != nil {
		return false, err
	}

	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	json.Unmarshal(data, &response)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *HTTPRequest) EraseAntrian(updateSpec businessTeknisi.UpdateJumlahAntrian) (bool, error) {
	var response UpdateAntrianResponse

	client := &http.Client{
		Timeout: time.Second * time.Duration(100),
	}

	bodyPayload := new(bytes.Buffer)
	if err := json.NewEncoder(bodyPayload).Encode(updateSpec); err != nil {
		log.Error(err)
		return false, nil
	}

	url := "http://127.0.0.1:7070/v1/teknisi/erase/antrian"
	request, err := http.NewRequest(http.MethodPut, url, bodyPayload)
	if err != nil {
		return false, err
	}

	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	json.Unmarshal(data, &response)

	if err != nil {
		return false, err
	}

	return true, nil
}
