package logger_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"task-api/src/util/clock"
	"task-api/src/util/logger"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	// bytes.Bufferを使ったwriterをモックする方法
	buffer := bytes.Buffer{}
	logger.Writer = &buffer

	logger.Debug("Hello World.")

	expected := "[DEBUG] Hello World.\n"
	actual := buffer.String()

	if expected != actual {
		t.Errorf("wont %s, but got %s", expected, actual)
	}
}

func TestDebugf(t *testing.T) {
	// os.Pipeを使ったwriterをモックする方法
	reader, writer, _ := os.Pipe()
	logger.Writer = writer

	logger.Debugf("Hello %s\n", "World.")
	writer.Close()

	expected := "[DEBUG] Hello World.\n"
	actual, _ := ioutil.ReadAll(reader)

	if expected != string(actual) {
		t.Errorf("wont %s, but got %s", expected, actual)
	}
}

func TestLog(t *testing.T) {
	buffer := bytes.Buffer{}

	logger.Writer = &buffer
	jst, _ := time.LoadLocation("Asia/Tokyo")
	clock.SetFakeTime(time.Date(2000, time.January, 01, 10, 10, 10, 10, jst))

	logger.Log("Hello World.")

	expected := "2000-01-01 10:10:10 Hello World.\n"
	actual := buffer.String()

	if expected != actual {
		t.Errorf("wont %s, but got %s", expected, actual)
	}
}
