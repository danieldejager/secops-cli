package wrappers

import (
	"encoding/json"
	"net/http"

	"github.com/danieldejager/secops-cli/internal/params"
	"github.com/spf13/viper"
)

const VT_API_URL = "https://www.virustotal.com/api/v3/files/"

var client *http.Client

type VirusTotalHTTPWrapper struct {
	path string
}

func NewVirusTotalWrapper(path string) VirusTotalWrapper {
	return &VirusTotalHTTPWrapper{path: path}
}

func (v *VirusTotalHTTPWrapper) GetFileAnalysis(SHA256 string) (VirusTotalResponse, error) {

	req, _ := http.NewRequest("GET", VT_API_URL+SHA256, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-apikey", viper.GetString(params.VirusTotalAPIKey))

	res, err := http.DefaultClient.Do(req)

	return handleResponse(res, err)

}

func handleResponse(resp *http.Response, err error) (VirusTotalResponse, error) {
	if err != nil {
		return VirusTotalResponse{}, err
	}

	decoder := json.NewDecoder(resp.Body)

	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {

	case http.StatusOK:
		err = decoder.Decode(VirusTotalResponse{})
		if err != nil {
			return VirusTotalResponse{}, err
		}
		return VirusTotalResponse{}, nil
	default:
		return VirusTotalResponse{}, nil
	}
}
