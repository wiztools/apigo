package apigo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func LogResp(r *http.Response) {
	LogRespErr(r, nil)
}

func LogRespErr(r *http.Response, err error) {
	if err == nil && r == nil {
		return
	}
	if err != nil && r != nil {
		log.Println("[api-resp-status-err]", r.Status, "|", err)
	} else if err != nil && r == nil {
		log.Println("[api-resp-err]", err)
	} else if err == nil && r != nil {
		log.Println("[api-resp-status]", r.Status)
	}

	// Body print:
	if r == nil {
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

func BindResp(r *http.Response, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
