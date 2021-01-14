### GoLogger
A simple logging utility for Go. Created to help learn the Go language.

### Dependencies
- [color](https://github.com/fatih/color)

### Example Usage
```go
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
    logger.Debug("REACHED THIS POINT")
    logger.Warning("REACHED THIS POINT")
    logger.Error("REACHED THIS POINT")
}
```
This will output the following:

> <div style="color: cyan;">[DEBUG][function][date | time] REACHED THIS POINT</div>
> <div style="color: yellow;">[WARNING][function][date | time] REACHED THIS POINT</div>
> <div style="color: red;">[ERROR][function][date | time] REACHED THIS POINT</div>

Each flag set in the instantiation of the Logger struct has a special influence on the output:
|Flag|Description|
|-|-|
|ShouldLogTime|Should any time information be logged?
|ShouldLogDate|If time information is logged, should that information include the date?
|ShouldLogFunc|Should the name of the calling function be logged?|
