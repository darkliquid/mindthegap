// +build debug

package engine

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stderr)
}

// Log logs output using the standard logger
func Log(str ...interface{}) {
	log.Print(str...)
}
