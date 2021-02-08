package b

import (
	. "github.com/maxinehazel/chroma" // nolint
	"github.com/maxinehazel/chroma/lexers/internal"
)

// Bnf lexer.
var Bnf = internal.Register(MustNewLexer(
	&Config{
		Name:      "BNF",
		Aliases:   []string{"bnf"},
		Filenames: []string{"*.bnf"},
		MimeTypes: []string{"text/x-bnf"},
	},
	Rules{
		"root": {
			{`(<)([ -;=?-~]+)(>)`, ByGroups(Punctuation, NameClass, Punctuation), nil},
			{`::=`, Operator, nil},
			{`[^<>:]+`, Text, nil},
			{`.`, Text, nil},
		},
	},
))
