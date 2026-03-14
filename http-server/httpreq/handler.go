package httpreq

import (
	"encoding/json"
	"net/http"
)

type StructValidator interface {
	Struct(s any) error
}

func DecodeBody(req *http.Request, body *any) error {
	return json.NewDecoder(req.Body).Decode(body)
}

func Validate(body any, validator StructValidator) error {
	return validator.Struct(body)
}
