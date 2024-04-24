package g1gateway

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"example.com/m/v2/room"
	"example.com/m/v2/utils"
)

type StatusData struct {
	MacAdress  string
	NrOfPeople int64
}

func (s *StatusData) Mac() string {
	return s.MacAdress
}

func (s *StatusData) Nr_of_people() int64 {
	return s.NrOfPeople
}

func ParseStatus(w http.ResponseWriter, r *http.Request) ([]StatusData, error) {
	data := make([]map[string]interface{}, 0)
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		slog.ErrorContext(r.Context(), "Failed to decode request body", "error", err)
		return nil, errors.New("request body is malformed")
	}

	if len(data) == 0 {
		return make([]StatusData, 0), nil
	}

	rooms, err := room.AllRooms()
	if err != nil {
		utils.WriteHttpError(w, "Internal server error", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Failed to get all rooms", "error", err)
		return nil, errors.New("failed to get all rooms")
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

		if macString == "" {
			slog.ErrorContext(r.Context(), "Mac address is empty", "mac", macString)
			continue
		}

		roomFound := false
		for _, room := range rooms {
			if room.Sensor == macString {
				roomFound = true
				break
			}
		}

		if !roomFound {
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

		if bytes[8] != 0x18 {
			continue
		}

		nr := int64((bytes[11] & 0x0F))

		parsedData = append(parsedData, StatusData{
			MacAdress:  macString,
			NrOfPeople: nr,
		})
	}

	return parsedData, nil
}
