package log

import "testing"

func before()  {
	Init()
}

func TestInfo(t *testing.T) {
	before()
	Info("ok")
}
