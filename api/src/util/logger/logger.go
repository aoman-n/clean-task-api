package logger

import (
	"fmt"
	"io"
	"os"
	"task-api/src/util/clock"
)

var Writer io.Writer = os.Stdout

func Debug(a ...interface{}) {
	a = append([]interface{}{"[DEBUG]"}, a...)
	fmt.Fprintln(Writer, a...)
}

func Debugf(format string, a ...interface{}) {
	fmt.Fprintf(Writer, "[DEBUG] "+format, a...)
}

func Log(a ...interface{}) {
	now := clock.Now().Format("2006-01-02 15:05:05")
	a = append([]interface{}{now}, a...)
	fmt.Fprintln(Writer, a...)
}
