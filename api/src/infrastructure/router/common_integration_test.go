// +build integration

package router

import (
	"os"
	"task-api/src/testutil"
	"testing"
)

func TestMain(m *testing.M) {
	testutil.SetUpDB("../../../db")
	code := m.Run()
	testutil.ClearDB()
	os.Exit(code)
}
