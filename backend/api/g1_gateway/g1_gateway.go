package g1gateway

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"example.com/m/v2/sensor"
	"example.com/m/v2/utils"
)

const (
	DATA_BYTE          int  = 11
	FRAME_VERSION_BYTE int  = 8
	FRAME_VERSION      byte = 0x18
	DATA_MASK          byte = 0x0F
)

type StatusData struct {
	MacAdress  string
	NrOfPeople int64
}

func ParseStatus(w http.ResponseWriter, r *http.Request) ([]StatusData, error) {
	data := make([]map[string]interface{}, 0)
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.WriteHttpError(w, "Request body is malformed", http.StatusBadRequest)
		slog.ErrorContext(r.Context(), "Failed to decode request body", "error", err)
		return nil, errors.New("request body is malformed")
	}

	if len(data) == 0 {
		return make([]StatusData, 0), nil
	}

	parsedData := make([]StatusData, 0)
	for _, dataItem := range data {
		mac, ok := dataItem["mac"]
		if !ok {
			continue
		}

		macString, ok := mac.(string)
		if !ok {
			continue
		}

		if _, err := sensor.RoomFromMac(macString); err != nil {
			continue
		}

		rawData, ok := dataItem["rawData"]
		if !ok {
			continue
		}

		stringData := rawData.(string)
		if rawData == "" {
			continue
		}

		bytes, err := hex.DecodeString(stringData)
		if err != nil {
			slog.ErrorContext(r.Context(), "Failed to decode hex string", "error", err)
			return nil, errors.New("failed to decode hex string")
		}

		if bytes[FRAME_VERSION_BYTE] != FRAME_VERSION {
			continue
		}

		nr := int64((bytes[DATA_BYTE] & DATA_MASK))

		parsedData = append(parsedData, StatusData{
			MacAdress:  macString,
			NrOfPeople: nr,
		})
	}

	return parsedData, nil
}
