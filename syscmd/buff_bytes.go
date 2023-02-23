//go:build !windows

package syscmd

func (lb *limitedBuffer) Bytes() []byte {
	return lb.bytes()
}
