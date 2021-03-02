package syntax

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

// Token it is recognized character group. Used by the parser.
type Token struct {
	// symbol understandable to parser
	Symbol int
	// start and end of data entering the symbol
	Start, End int

	// name of the group
	Labels map[string]bool

	// file position
	Pos Pos
}

func (t *Token) AddLabels(labels ...string) {
	if t.Labels == nil {
		t.Labels = map[string]bool{}
	}
	for _, label := range labels {
		t.Labels[label] = true
	}
}

func (t *Token) MatchAtLeastOneLabels(labels ...string) bool {
	for _, label := range labels {
		if t.Labels[label] {
			return true
		}
	}
	return false
}

type Tokens []*Token

func (t *Tokens) Add(in ...*Token) int {
	*t = append(*t, in...)
	return len(*t) - 1
}

type DataBytes []byte

func (dat DataBytes) Render(t *Token) ([]byte, error) {
	r := bytes.NewReader(dat)

	_, err := r.Seek(int64(t.Start), io.SeekStart)
	if err != nil {
		return nil, errors.Wrap(err, "failed seekeing")
	}
	res := make([]byte, t.End-t.Start)

	_, err = r.Read(res)
	if err != nil {
		return nil, errors.Wrap(err, "failed read")
	}

	return res, nil
}
