# Goroutine count

A small library to get the number of routines spawned from a specific line of code
and are in a given state.

# API

```go
package main

import (
	"github.com/tomzx/goroutine-count/pkg/counter"
)

//type Goroutine struct {
//    Identifier string // e.g., "/path/to/file.go:123"
//    State      string // e.g., "running", "runnable", "sleep", etc.
//}

// Returns a map[Goroutine]int, where the value is the number of routines
// for the given Goroutine (identifier and state).
counts := counter.GetGoroutineCount()
```

## License

The code is licensed under the [MIT license](http://choosealicense.com/licenses/mit/). See [LICENSE](LICENSE).