package js

import "fmt"

var (
	Green  = Color("\033[1;32m%s\033[0m")
	Red    = Color("\033[1;31m%s\033[0m")
	Teal   = Color("\033[1;36m%s\033[0m")
	White  = Color("\033[1;37m%s\033[0m")
	Yellow = Color("\033[1;33m%s\033[0m")
)

var (
	Info = Teal
	Warn = Yellow
	Err  = Red
)

func Color(colorString string) func(string, ...interface{}) string {
	return func(format string, args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprintf(format, args...))
	}
}
