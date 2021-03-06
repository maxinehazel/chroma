package i

import (
	. "github.com/maxinehazel/chroma" // nolint
	"github.com/maxinehazel/chroma/lexers/internal"
)

// Io lexer.
var Io = internal.Register(MustNewLexer(
	&Config{
		Name:      "Io",
		Aliases:   []string{"io"},
		Filenames: []string{"*.io"},
		MimeTypes: []string{"text/x-iosrc"},
	},
	Rules{
		"root": {
			{`\n`, Text, nil},
			{`\s+`, Text, nil},
			{`//(.*?)\n`, CommentSingle, nil},
			{`#(.*?)\n`, CommentSingle, nil},
			{`/(\\\n)?[*](.|\n)*?[*](\\\n)?/`, CommentMultiline, nil},
			{`/\+`, CommentMultiline, Push("nestedcomment")},
			{`"(\\\\|\\"|[^"])*"`, LiteralString, nil},
			{`::=|:=|=|\(|\)|;|,|\*|-|\+|>|<|@|!|/|\||\^|\.|%|&|\[|\]|\{|\}`, Operator, nil},
			{`(clone|do|doFile|doString|method|for|if|else|elseif|then)\b`, Keyword, nil},
			{`(nil|false|true)\b`, NameConstant, nil},
			{`(Object|list|List|Map|args|Sequence|Coroutine|File)\b`, NameBuiltin, nil},
			{`[a-zA-Z_]\w*`, Name, nil},
			{`(\d+\.?\d*|\d*\.\d+)([eE][+-]?[0-9]+)?`, LiteralNumberFloat, nil},
			{`\d+`, LiteralNumberInteger, nil},
		},
		"nestedcomment": {
			{`[^+/]+`, CommentMultiline, nil},
			{`/\+`, CommentMultiline, Push()},
			{`\+/`, CommentMultiline, Pop(1)},
			{`[+/]`, CommentMultiline, nil},
		},
	},
))
