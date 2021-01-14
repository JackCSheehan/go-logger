package main


import (
	"./goLogger"
)

// Example goLogger usage
func main() {
	logger := goLogger.Logger{true, true, true}
	logger.Debug("REACHED THIS POINT")
}