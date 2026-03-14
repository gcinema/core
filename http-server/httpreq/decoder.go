package httpreq

import (
	"encoding/json"
	"io"
)

func decodeBody[T any](body io.ReadCloser) (*T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
