package main

import (
	"os"
	"time"

	"example.com/me/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
