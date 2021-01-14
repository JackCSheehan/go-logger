package main


import (
	"./goLogging"
)

func main() {
	logger := goLogging.Logger{true, false}
	logger.Debug("Hello, World!")
}