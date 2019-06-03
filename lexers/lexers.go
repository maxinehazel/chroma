// Package lexers contains the registry of all lexers.
//
// Sub-packages contain lexer implementations.
package lexers

// nolint: golint
import (
	"github.com/softpunks/chroma"
	_ "github.com/softpunks/chroma/lexers/a"
	_ "github.com/softpunks/chroma/lexers/b"
	_ "github.com/softpunks/chroma/lexers/c"
	_ "github.com/softpunks/chroma/lexers/circular"
	_ "github.com/softpunks/chroma/lexers/d"
	_ "github.com/softpunks/chroma/lexers/e"
	_ "github.com/softpunks/chroma/lexers/f"
	_ "github.com/softpunks/chroma/lexers/g"
	_ "github.com/softpunks/chroma/lexers/h"
	_ "github.com/softpunks/chroma/lexers/i"
	"github.com/softpunks/chroma/lexers/internal"
	_ "github.com/softpunks/chroma/lexers/j"
	_ "github.com/softpunks/chroma/lexers/k"
	_ "github.com/softpunks/chroma/lexers/l"
	_ "github.com/softpunks/chroma/lexers/m"
	_ "github.com/softpunks/chroma/lexers/n"
	_ "github.com/softpunks/chroma/lexers/o"
	_ "github.com/softpunks/chroma/lexers/p"
	_ "github.com/softpunks/chroma/lexers/q"
	_ "github.com/softpunks/chroma/lexers/r"
	_ "github.com/softpunks/chroma/lexers/s"
	_ "github.com/softpunks/chroma/lexers/t"
	_ "github.com/softpunks/chroma/lexers/v"
	_ "github.com/softpunks/chroma/lexers/w"
	_ "github.com/softpunks/chroma/lexers/x"
	_ "github.com/softpunks/chroma/lexers/y"
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
