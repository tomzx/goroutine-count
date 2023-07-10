package counter

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetGoroutineCount(t *testing.T) {
	_, filename, line, _ := runtime.Caller(0)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			time.Sleep(5 * time.Second)
		}(i)
	}

	actual := GetGoroutineCount()
	expected := map[Goroutine]int{
		{Identifier: fmt.Sprintf("%s:%d", filename, line+2), State: "sleep"}: 10000,
		{Identifier: "_testmain.go:49", State: "chan receive"}:               1,
	}
	// Remove testing/testing.go entry as it is not consistent across different platforms
	for routine, _ := range actual {
		if strings.Contains(routine.Identifier, "testing/testing.go") {
			delete(actual, routine)
		}
	}
	assert.Equal(t, expected, actual)
}

func BenchmarkGetGoroutineCount(b *testing.B) {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			time.Sleep(5 * time.Second)
		}(i)
	}
	for i := 0; i < b.N; i++ {
		GetGoroutineCount()
	}
}
