package Use

import (
	"ch11/Series"
	"testing"
)

func TestFib(t *testing.T) {
	t.Log(Series.FibonacciGrt(5))
	t.Log(Series.Square(5))

}
