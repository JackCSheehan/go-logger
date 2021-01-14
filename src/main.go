package main


import (
	"./goLogging"
)

func main() {
	logger := goLogging.Logger{true, false}
	logger.Debug("Here's a debug message")
	logger.Warning("Here's a warning message")
	logger.Error("Here's an error message")
}