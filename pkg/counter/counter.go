package counter

import (
	"regexp"
	"runtime"
	"strings"
)

type Goroutine struct {
	Identifier string
	State      string
}

func getGoroutines() []Goroutine {
	buf := make([]byte, 1<<24) // 16 MB
	length := runtime.Stack(buf, true)
	stack := string(buf[:length])

	stateRegex := regexp.MustCompile("goroutine \\d+ \\[(.*)\\]")
	identifierRegex := regexp.MustCompile("\\t(.*) \\+0x")
	identifiers := make([]Goroutine, 0, 1024)

	identifier := ""
	state := ""
	lines := strings.Split(stack, "\n")
	for i, line := range lines {
		// Capture the state of the goroutine
		if strings.HasPrefix(line, "goroutine") {
			matches := stateRegex.FindStringSubmatch(line)
			if len(matches) == 0 {
				break
			}
			state = matches[1]
			continue
		}

		// If we found an empty line, this indicates we're between Goroutines stacks
		// and that the previous line contains the identifier we want to capture.
		if line == "" {
			previousLine := lines[i-1]
			matches := identifierRegex.FindStringSubmatch(previousLine)
			if len(matches) == 0 {
				break
			}
			identifier = matches[1]
			identifiers = append(identifiers, Goroutine{Identifier: identifier, State: state})
			state = ""
			identifier = ""
			continue
		}
	}

	return identifiers
}

func GetGoroutineCount() map[Goroutine]int {
	goroutines := getGoroutines()
	counts := make(map[Goroutine]int)
	for _, goroutine := range goroutines {
		counts[goroutine]++
	}
	return counts
}
