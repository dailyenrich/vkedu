package log

import (
	"fmt"
	"testing"
	"time"
)

func before()  {
	Init(fmt.Sprintf(
		"runtime/log/%s.log",
		time.Now().Format(fileNameFormat),
	))
}

func TestInfo(t *testing.T) {
	before()
	Info("ok")
}
