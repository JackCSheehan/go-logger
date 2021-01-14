/*
File containing defintion of Logger struct
*/
package goLogging

import (
	"fmt"
	"time"
	"github.com/fatih/color"
)

// Color constants for each output type
const debugColor = color.FgHiCyan
const warningColor = color.FgHiYellow
const errorColor = color.FgHiRed

type LogType int

// Logger output types
const (
	debug LogType = iota	// Type for basic debug logging
	warn					// Type for warning logs
	err						// Type for error logs
)

/*
Struct containing methods needed for logging
*/
type Logger struct {
	ShouldLogTime	bool	// If time should be logged
	ShouldLogDate	bool	// If date should be logged with time
}

/*
Method to print the current time
*/
func (logger Logger) logTime() {
	// Get current time
	nowTime := time.Now()

	// Print opening bracket
	fmt.Print("[")

	// Print date if date flag set
	if logger.ShouldLogDate {
		fmt.Print(nowTime.Date())
	}

	// Print time
	fmt.Printf("%02d:%02d:%02d]", nowTime.Hour(), nowTime.Minute(), nowTime.Second())
}

/*
Private helper method to handle logging based on a given LogType. Also takes the
message to be printed
*/
func (logger Logger) log(logType LogType, message string) {
	// Change color and print indicator based on log type
	if logType == debug {
		color.Set(debugColor)
		fmt.Print("DEBUG:")
	} else if logType == warn {
		color.Set(warningColor)
		fmt.Print("WARNING:")
	} else {
		color.Set(errorColor)
		fmt.Print("ERROR:")
	}

	// Log time if time flag set
	if logger.ShouldLogTime {
		logger.logTime()
	}

	// Print log message
	fmt.Println(" ", message)

	color.Unset()
}

/*
Method for displaying a debug message
*/
func (logger Logger) Debug(message string) {
	logger.log(debug, message)
}

/*
Method for displaying a warning message
*/
func (logger Logger) Warning(message string) {
	logger.log(warn, message)
}

/*
Method for displaying an error message
*/
func (logger Logger) Error(message string) {
	logger.log(err, message)
}