package database

import (
	"os"
)

var FileLog *os.File

func AddStr(str string) {
	if FileLog != nil {
		_, _ = FileLog.WriteString("\n" + str)
	}
}
