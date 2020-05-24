package errors

type ModelValidationErr struct {
	Msg string
}

func (e *ModelValidationErr) Error() string {
	return "ValidationError: " + e.Msg
}

func NewModelValidationErr(msg string) *ModelValidationErr {
	return &ModelValidationErr{msg}
}

type RecordSaveErr struct {
	Msg string
}

func (e *RecordSaveErr) Error() string {
	return "SaveRecordError: " + e.Msg
}

func NewRecordSaveErr(msg string) *RecordSaveErr {
	return &RecordSaveErr{msg}
}

type PermissionErr struct {
	Msg string
}

func (e *PermissionErr) Error() string {
	return "PermissionError: " + e.Msg
}

func NewPermissionErr(msg string) *PermissionErr {
	return &PermissionErr{msg}
}

type NotFoundErr struct {
	Msg string
}

func (e *NotFoundErr) Error() string {
	return "NotFoundErr: " + e.Msg
}

func NewNotFoundErr(msg string) *NotFoundErr {
	return &NotFoundErr{msg}
}
