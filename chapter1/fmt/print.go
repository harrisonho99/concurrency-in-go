package fmt

import (
	"fmt"
)

const (
	InfoColor    = "\033[1;34m%v\033[0m"
	NoticeColor  = "\033[1;36m%v\033[0m"
	WarningColor = "\033[1;33m%v\033[0m"
	ErrorColor   = "\033[1;31m%v\033[0m"
	DebugColor   = "\033[0;36m%v\033[0m"
)

func Println(i ...interface{}) {
	fmt.Printf(InfoColor+"\n", i...)
}

func Printf(i ...interface{}) {
	fmt.Printf(InfoColor, fmt.Sprintf(i[0].(string), i[1:]...))
}
