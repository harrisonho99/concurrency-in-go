package fmt

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Color string

const (
	InfoColor    Color = "\033[1;34m%v\033[0m"
	NoticeColor  Color = "\033[1;36m%v\033[0m"
	WarningColor Color = "\033[1;33m%v\033[0m"
	ErrorColor   Color = "\033[1;31m%v\033[0m"
	DebugColor   Color = "\033[0;36m%v\033[0m"
)

var (
	defaultColorProfile Color = InfoColor
)

func SetColorProfile(c Color) {
	defaultColorProfile = c
}

func Fprintf(wr io.Writer, i ...interface{}) {
	fmt.Fprintf(wr, fmt.Sprintf(string(defaultColorProfile), i[0]), i[1:]...)
}
func Fprint(wr io.Writer, i ...interface{}) {
	var str strings.Builder
	for index := range i {
		str.WriteString(fmt.Sprintf("%v", i[index]))
		if index != len(i)-1 {
			// separator
			str.WriteString(" ")
		}
	}
	fmt.Fprintf(wr, string(defaultColorProfile), str.String())
}
func Printf(i ...interface{}) {
	Fprintf(os.Stdout, i...)
}
func Println(i ...interface{}) {
	Fprint(os.Stdout, i...)
	Fprint(os.Stdout, "\n")
}
