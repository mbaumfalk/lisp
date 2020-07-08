// +build !js

package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	NewInteractiveState().ReadFrom(os.Stdin)
}
