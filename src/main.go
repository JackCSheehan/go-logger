package main

import (
	"./goLogger"
)

// Example goLogger usage
func main() {
	logger := goLogger.Logger {
		ShouldLogTime: true,
		ShouldLogDate: true,
		ShouldLogFunc: true,
	}
	logger.Warning("REACHED THIS POINT")
}