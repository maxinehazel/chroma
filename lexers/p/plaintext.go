package p

import (
	. "github.com/maxinehazel/chroma" // nolint
	"github.com/maxinehazel/chroma/lexers/internal"
)

var Plaintext = internal.Register(MustNewLexer(
	&Config{
		Name:      "plaintext",
		Aliases:   []string{"text", "plain", "no-highlight"},
		Filenames: []string{"*.txt"},
		MimeTypes: []string{"text/plain"},
	},
	internal.PlaintextRules,
))
