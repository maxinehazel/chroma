package b

import (
	. "github.com/maxinehazel/chroma" // nolint
	"github.com/maxinehazel/chroma/lexers/internal"
)

// Brainfuck lexer.
var Brainfuck = internal.Register(MustNewLexer(
	&Config{
		Name:      "Brainfuck",
		Aliases:   []string{"brainfuck", "bf"},
		Filenames: []string{"*.bf", "*.b"},
		MimeTypes: []string{"application/x-brainfuck"},
	},
	Rules{
		"common": {
			{`[.,]+`, NameTag, nil},
			{`[+-]+`, NameBuiltin, nil},
			{`[<>]+`, NameVariable, nil},
			{`[^.,+\-<>\[\]]+`, Comment, nil},
		},
		"root": {
			{`\[`, Keyword, Push("loop")},
			{`\]`, Error, nil},
			Include("common"),
		},
		"loop": {
			{`\[`, Keyword, Push()},
			{`\]`, Keyword, Pop(1)},
			Include("common"),
		},
	},
))
