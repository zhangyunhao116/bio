// +build !go1.16

// Package xio holds the transition packages for the new Go 1.16 io.ReadAll etc.
package xio

import (
	"io"
	"io/ioutil"
)

func ReadAll(r io.Reader) ([]byte,error) {
	return  ioutil.ReadAll(r)
}