package main

import (
	"os"
	"time"

	clockface "github.com/gomesmf/go-learning/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
