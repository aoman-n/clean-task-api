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
