// Package lexers contains the registry of all lexers.
//
// Sub-packages contain lexer implementations.
package lexers

// nolint: golint
import (
	"github.com/maxinehazel/chroma"
	_ "github.com/maxinehazel/chroma/lexers/a"
	_ "github.com/maxinehazel/chroma/lexers/b"
	_ "github.com/maxinehazel/chroma/lexers/c"
	_ "github.com/maxinehazel/chroma/lexers/circular"
	_ "github.com/maxinehazel/chroma/lexers/d"
	_ "github.com/maxinehazel/chroma/lexers/e"
	_ "github.com/maxinehazel/chroma/lexers/f"
	_ "github.com/maxinehazel/chroma/lexers/g"
	_ "github.com/maxinehazel/chroma/lexers/h"
	_ "github.com/maxinehazel/chroma/lexers/i"
	"github.com/maxinehazel/chroma/lexers/internal"
	_ "github.com/maxinehazel/chroma/lexers/j"
	_ "github.com/maxinehazel/chroma/lexers/k"
	_ "github.com/maxinehazel/chroma/lexers/l"
	_ "github.com/maxinehazel/chroma/lexers/m"
	_ "github.com/maxinehazel/chroma/lexers/n"
	_ "github.com/maxinehazel/chroma/lexers/o"
	_ "github.com/maxinehazel/chroma/lexers/p"
	_ "github.com/maxinehazel/chroma/lexers/q"
	_ "github.com/maxinehazel/chroma/lexers/r"
	_ "github.com/maxinehazel/chroma/lexers/s"
	_ "github.com/maxinehazel/chroma/lexers/t"
	_ "github.com/maxinehazel/chroma/lexers/v"
	_ "github.com/maxinehazel/chroma/lexers/w"
	_ "github.com/maxinehazel/chroma/lexers/x"
	_ "github.com/maxinehazel/chroma/lexers/y"
)

// Registry of Lexers.
var Registry = internal.Registry

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string { return internal.Names(withAliases) }

// Get a Lexer by name, alias or file extension.
func Get(name string) chroma.Lexer { return internal.Get(name) }

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) chroma.Lexer { return internal.MatchMimeType(mimeType) }

// Match returns the first lexer matching filename.
func Match(filename string) chroma.Lexer { return internal.Match(filename) }

// Analyse text content and return the "best" lexer..
func Analyse(text string) chroma.Lexer { return internal.Analyse(text) }

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer) chroma.Lexer { return internal.Register(lexer) }

// Fallback lexer if no other is found.
var Fallback = internal.Fallback
