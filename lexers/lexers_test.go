package lexers_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/maxinehazel/assert"
	"github.com/maxinehazel/chroma"
	"github.com/maxinehazel/chroma/formatters"
	"github.com/maxinehazel/chroma/lexers"
	"github.com/maxinehazel/chroma/lexers/a"
	"github.com/maxinehazel/chroma/lexers/x"
	"github.com/maxinehazel/chroma/styles"
)

func TestCompileAllRegexes(t *testing.T) {
	for _, lexer := range lexers.Registry.Lexers {
		it, err := lexer.Tokenise(nil, "")
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
		err = formatters.NoOp.Format(ioutil.Discard, styles.SwapOff, it)
		assert.NoError(t, err, "%s failed", lexer.Config().Name)
	}
}

func TestGet(t *testing.T) {
	t.Run("ByName", func(t *testing.T) {
		assert.Equal(t, lexers.Get("xml"), x.XML)
	})
	t.Run("ByAlias", func(t *testing.T) {
		assert.Equal(t, lexers.Get("as"), a.Actionscript)
	})
	t.Run("ViaFilename", func(t *testing.T) {
		assert.Equal(t, lexers.Get("svg"), x.XML)
	})
}

// Test source files are in the form <key>.<key> and validation data is in the form <key>.<key>.expected.
func TestLexers(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	assert.NoError(t, err)

	for _, file := range files {
		ext := filepath.Ext(file.Name())[1:]
		if ext != "actual" {
			continue
		}

		lexer := lexers.Get(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
		assert.NotNil(t, lexer)

		filename := filepath.Join("testdata", file.Name())
		expectedFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".expected"

		lexer = chroma.Coalesce(lexer)
		t.Run(lexer.Config().Name, func(t *testing.T) {
			// Read and tokenise source text.
			actualText, err := ioutil.ReadFile(filename)
			assert.NoError(t, err)
			actual, err := chroma.Tokenise(lexer, nil, string(actualText))
			assert.NoError(t, err)

			if os.Getenv("RECORD") == "true" {
				// Update the expected file with the generated output of this lexer
				f, err := os.Create(expectedFilename)
				defer f.Close()
				assert.NoError(t, err)
				assert.NoError(t, formatters.JSON.Format(f, nil, chroma.Literator(actual...)))
			} else {
				// Read expected JSON into token slice.
				var expected []chroma.Token
				r, err := os.Open(expectedFilename)
				assert.NoError(t, err)
				err = json.NewDecoder(r).Decode(&expected)
				assert.NoError(t, err)

				// Equal?
				assert.Equal(t, expected, actual)
			}
		})
	}
}
