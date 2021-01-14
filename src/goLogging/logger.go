/*
File containing defintion of Logger struct
*/
package goLogging

import (
	"fmt"
	"time"
	"github.com/fatih/color"
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
func (logger Logger) LogTime() {
	// Get current time
	nowTime := time.Now()

	// Print opening bracket
	fmt.Print("[")

	// Print date if date flag set
	if logger.ShouldLogDate {
		fmt.Print(nowTime.Date())
	}

	// Print time
	fmt.Printf("%d:%d:%02d]", nowTime.Hour(), nowTime.Minute(), nowTime.Second())
}

/*
Prints Debug message
*/
func (logger Logger) Debug(message string) {
	color.Set(color.FgYellow)

	// Print debug indicator
	fmt.Print("DEBUG:")

	// Log time if time flag set
	if logger.ShouldLogTime {
		logger.LogTime()
	}

	// Print debug message
	fmt.Println(" ", message)

	color.Unset()
}