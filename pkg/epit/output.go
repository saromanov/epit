package epit

import "github.com/fatih/color"

// info retruns text for information
func info(format string, a ...interface{}) {
	color.Blue(format, a...)
}

// fail returns text for error
func fail(format string, a ...interface{}) {
	color.Red(format, a...)
}
