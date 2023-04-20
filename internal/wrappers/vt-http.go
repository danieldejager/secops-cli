package wrappers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danieldejager/secops-cli/internal/params"
	"github.com/spf13/viper"
)

const (
	failedToGetFileObject = "failed to parse get request"
)

type VirusTotalHTTPWrapper struct {
	path string
}

func NewHTTPVirusTotalWrapper(path string) VirusTotalWrapper {
	return &VirusTotalHTTPWrapper{path: path}
}

func (v *VirusTotalHTTPWrapper) GetFileAnalysis(hash string) (
	*VirusTotalResponseModel,
	error) {
	resp, err := http.NewRequest(http.MethodGet, viper.GetString(params.VTBasePathEnv+params.VTFilePathEnv), nil)
	if err != nil {
		log.Fatal("Cannot request File Report from VT")
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)

	defer resp.Body.Close()

	switch resp.Response.StatusCode {
	case http.StatusOK:
		model := VirusTotalResponseModel{}
		err = decoder.Decode(&model)
		if err != nil {
			return nil, nil
		}
		return &model, nil

	case http.StatusBadRequest, http.StatusInternalServerError:
		return nil, nil

	default:
		return nil, nil
	}
}
