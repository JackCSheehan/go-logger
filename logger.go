/*
File containing defintion of Logger struct
*/
package goLogger

import (
	"fmt"
	"time"
	"github.com/fatih/color"
	"runtime"
	"strconv"
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
	ShouldLogFunc	bool	// If current function should be logged
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
		// Format year and month for printing
		year := strconv.Itoa(nowTime.Year())[2:]
		day := strconv.Itoa(nowTime.Day())
		month := nowTime.Month().String()[:3]

		fmt.Printf("%s %s '%s | ", month, day, year)
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
		fmt.Print("[DEBUG]")
	} else if logType == warn {
		color.Set(warningColor)
		fmt.Print("[WARNING]")
	} else {
		color.Set(errorColor)
		fmt.Print("[ERROR]")
	}

	// Log current function name if function flag set
	if logger.ShouldLogFunc {
		fmt.Printf("[%s]", logger.getFunctionName())
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

/*
Private helper function which returns the name of the function currently being run
*/
func (logger Logger) getFunctionName() string {
	// Get name of current function
	pc, _, _, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc)

	return fn.Name()
}