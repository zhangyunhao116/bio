// Package bio provides better io utils for Go.
package bio

import (
	"io"
)

// FixedReadAll reads from r into a preallocated slice until an error or EOF.
// If size <= 0 or size > limit, return io.ReadAll(r) directly.
// If limit <= 0 means the preallocated slice is unlimited.
// Even if the actual data amount exceeds the size, this function still read all data from r.
func FixedReadAll(r io.Reader, size int, limit int) ([]byte, error) {
	if size <= 0 || (limit > 0 && size > limit) {
		return io.ReadAll(r)
	}
	data := make([]byte, size+1)

	var i int
	for {
		n, err := r.Read(data[i:])
		if n < 0 {
			panic("reader returned negative count from Read")
		}
		if err == io.EOF { // err is EOF, so return nil explicitly
			i += n
			break
		} else if err != nil {
			return data, err
		}
		i += n
		if i == size+1 { // actual data amount exceeds the size
			rest, err := io.ReadAll(r)
			res := make([]byte, len(data)+len(rest))
			copy(res[:len(data)], data)
			copy(res[len(data):], rest)
			return res, err
		}
	}
	return data[:i], nil
}
