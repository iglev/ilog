package log

import "testing"

func TestLog(t *testing.T) {
	Info("num=%v", 1234)
}
