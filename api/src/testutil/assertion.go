package testutil

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertResponse レスポンスのHeaderとBodyを検証
func AssertResponse(t *testing.T, res *httptest.ResponseRecorder, code int, expectedJSONath string) {
	t.Helper()

	AssertResponseHeader(t, res, code)
	AssertResponseBodyWithFile(t, res, expectedJSONath)
}

// AssertResponseHeader レスポンスヘッダーのStatusCodeとContent-Typeを検証
func AssertResponseHeader(t *testing.T, res *httptest.ResponseRecorder, code int) {
	t.Helper()

	// ステータスコードのチェック
	if code != res.Code {
		t.Errorf("expected status code is '%d',\n but actual given code is '%d'", code, res.Code)
	}
	// Content-Typeのチェック
	if expected := "application/json"; res.Header().Get("Content-Type") != expected {
		t.Errorf("unexpected response Content-Type,\n expected: %#v,\n but given #%v", expected, res.Header().Get("Content-Type"))
	}
}

// AssertResponseBodyWithFile レスポンスボディのJSONを検証
func AssertResponseBodyWithFile(t *testing.T, res *httptest.ResponseRecorder, path string) {
	t.Helper()

	expectedJSON := GetStringFromTestFile(t, path)
	assert.JSONEq(t, expectedJSON, res.Body.String())
}
