package epit

import "github.com/fatih/color"

// info retruns text for information
func info(format string, a ...interface{}) {
	color.Blue(format, a...)
}
