// +build appengine

package ruslog

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
