package httpreq

type StructValidator interface {
	Struct(s any) error
}

func validate[T any](validator StructValidator, payload *T) error {
	return validator.Struct(payload)
}
