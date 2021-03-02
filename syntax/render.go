package syntax

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

func NewRender(dat []byte) *TokenRender {
	return &TokenRender{
		dat: bytes.NewReader(dat),
	}
}

type TokenRender struct {
	dat interface {
		io.Reader
		io.Seeker
	}
}

func (l *TokenRender) Render(t *Token) ([]byte, error) {
	// defer l.dat.Seek(0, io.SeekStart)

	_, err := l.dat.Seek(int64(t.Start), io.SeekStart)
	if err != nil {
		return nil, errors.Wrap(err, "failed seekeing")
	}
	res := make([]byte, t.End-t.Start+1)

	_, err = l.dat.Read(res)
	if err != nil {
		return nil, errors.Wrap(err, "failed read")
	}

	return res, nil
}
