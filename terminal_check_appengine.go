// +build appengine

package shuzilm

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
