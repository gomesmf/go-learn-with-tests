package main

import (
	"os"
	"time"

	clockface "github.com/gomesmf/go-learn-with-tests/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
