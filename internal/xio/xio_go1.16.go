// +build go1.16

package xio

import (
	"io"
)

func ReadAll(r io.Reader) ([]byte,error) {
	return  io.ReadAll(r)
}