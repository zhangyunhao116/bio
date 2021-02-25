package bio

import (
	"fmt"
	"io"
	"testing"

	"github.com/zhangyunhao116/bio/internal/xio"
)

type mockReader struct {
	size int
	now  int
	step int
}

func newMockReader(size int, step int) *mockReader {
	return &mockReader{size: size, now: 0, step: step}
}

func (m *mockReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	step := m.step
	if len(p) < m.step {
		step = len(p)
	}

	var i int
	for i < step {
		p[i] = byte(m.now)
		m.now++
		if m.now == 256 {
			m.now = 0
		}
		m.size--
		i++
		if m.size == 0 {
			return i, io.EOF
		}
	}
	return i, nil
}

func bytesEqual(b1 []byte, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i, v := range b1 {
		if b2[i] != v {
			return false
		}
	}
	return true
}

func TestAll(t *testing.T) {
	const length = 1000

	d2, err := xio.ReadAll(newMockReader(length, 1))
	if err != nil {
		t.Fatalf("%s", err)
	}

	// All limit and size.
	for _, limit := range []int64{-100, -1, 0, 1, 32, length - 1, length, length + 1, length + 1024} {
		for _, size := range []int64{-100, -1, 0, 1, 32, length - 1, length, length + 1, length + 1024} {
			d1, err := FixedReadAll(newMockReader(length, 1), size, limit)
			if err != nil {
				t.Fatalf("%s", err)
			}
			if !bytesEqual(d1, d2) {
				t.Fatalf("unequal size:%d limit:%d\r\n`%s`\r\n`%s`", size, limit, string(d1), string(d2))
			}
		}
	}
}

func BenchmarkAll(b *testing.B) {
	const MB = 1024 * 1024
	for _, length := range []int{1 * MB, 5 * MB, 10 * MB, 100 * MB, 1024 * MB} {
		b.Run(fmt.Sprintf("FixedReadAll- %d MB", length/MB), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FixedReadAll(newMockReader(length, 100), int64(length), -1) //nolint:errcheck
			}
			b.ReportAllocs()
		})
		b.Run(fmt.Sprintf("io.ReadAll-   %d MB", length/MB), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				xio.ReadAll(newMockReader(length, 100)) //nolint:errcheck
			}
			b.ReportAllocs()
		})
	}
}
