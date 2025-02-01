package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Quan0308/main-api/core/dtos"
)

func Decode(r *http.Request, val interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(val); err != nil {
		return err
	}
	return nil
}

func Response(w http.ResponseWriter, resp dtos.Response) error {
	if resp.Message == "" {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			resp.Message = "OK"
		} else {
			resp.Message = "Error"
		}
	}

	res, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}
