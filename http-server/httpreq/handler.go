// Package httpreq contains HTTP request helpers.
package httpreq

import (
	"net/http"

	"github.com/gcinema/core/http-server/httpres"
)

func DecodeAndValidateBody[T any](w *http.ResponseWriter, req *http.Request, validator StructValidator) (*T, error) {
	payload, err := decodeBody[T](req.Body)
	if err != nil {
		httpres.ConvertToJSON(*w, err.Error(), 402)
		return nil, err
	}

	err = validate(validator, &payload)
	if err != nil {
		httpres.ConvertToJSON(*w, err.Error(), 402)
		return nil, err
	}

	return payload, nil
}
