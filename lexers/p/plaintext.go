package p

import (
	. "github.com/maxinekrebs/chroma" // nolint
	"github.com/maxinekrebs/chroma/lexers/internal"
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
