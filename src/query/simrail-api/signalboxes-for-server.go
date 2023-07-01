package simrail_api

import (
	"encoding/json"
	"errors"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types/simrail-api/responses"
	"io"
	"net/http"
)

func SignalboxesForServerQuery(shortName string) ([]responses.Signalbox, error) {
	client := http.Client{}
	url := "https://panel.simrail.eu:8084/stations-open?serverCode=" + shortName

	response, bodyReaderError := client.Get(url)
	if bodyReaderError != nil {
		return nil, errors.New("request failed to " + url)
	}

	defer response.Body.Close()

	body, bodyReaderError := io.ReadAll(response.Body)
	if bodyReaderError != nil {
		return nil, errors.New("BodyReader failed")
	}

	var responseBody responses.SignalboxesForServerResponse
	unmarshalError := json.Unmarshal(body, &responseBody)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return responseBody.Data, nil
}
