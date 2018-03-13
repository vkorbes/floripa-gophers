package main

import (
	"flag"
)

func main() {
	// START OMIT
	x := flag.Int("x", 0, "Essa explicação é\nbem longa, sacoé?")
	y := flag.Int("y", 0, "Essa é curtinha.")
	flag.Parse()
	// END OMIT
	flag.PrintDefaults()
	_, _ = x, y
}