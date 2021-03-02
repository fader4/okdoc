package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerRenderer(t *testing.T) {
	t.Run("", func(t *testing.T) {
		tok := &Token{Start: 0, End: 1}
		r := NewRender([]byte("abc"))
		res, err := r.Render(tok)
		assert.NoError(t, err)
		assert.EqualValues(t, string("ab"), string(res))
	})

	t.Run("", func(t *testing.T) {
		tok := &Token{Start: 0, End: 2}
		r := NewRender([]byte("abc"))
		res, err := r.Render(tok)
		assert.NoError(t, err)
		assert.EqualValues(t, string("abc"), string(res))
	})

	t.Run("", func(t *testing.T) {
		tok := &Token{Start: 1, End: 1}
		r := NewRender([]byte("abc"))
		res, err := r.Render(tok)
		assert.NoError(t, err)
		assert.EqualValues(t, string("b"), string(res))
	})

	t.Run("", func(t *testing.T) {
		tok := &Token{Start: 0, End: 0}
		r := NewRender([]byte("abc"))
		res, err := r.Render(tok)
		assert.NoError(t, err)
		assert.EqualValues(t, string("a"), string(res))
	})

}
