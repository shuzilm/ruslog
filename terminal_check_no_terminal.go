// +build js nacl plan9

package ruslog

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
