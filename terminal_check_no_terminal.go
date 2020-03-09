// +build js nacl plan9

package shuzilm

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
