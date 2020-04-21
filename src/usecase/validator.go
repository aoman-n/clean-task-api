package usecase

type Validator interface {
	Struct(interface{}) error
}
