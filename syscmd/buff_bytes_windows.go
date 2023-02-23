package syscmd

import "golang.org/x/text/encoding/simplifiedchinese"

func (lb *limitedBuffer) Bytes() []byte {
	buf := lb.bytes()
	if dec, err := simplifiedchinese.GB18030.NewDecoder().Bytes(buf); err == nil {
		return dec
	}
	return buf
}
