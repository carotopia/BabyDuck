package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/symbols"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// Parser wraps the ANTLR parser and directory builder
type Parser struct {
	sourceCode       string
	debug            bool
	directoryBuilder *DirectoryBuilder
	functionDir      *symbols.FunctionDirectory
	constantTable    *symbols.ConstantTable
}

// NewParser creates a new parser instance
func NewParser(sourceCode string, debug bool) *Parser {
	// Crear symbol tables sin memory manager
	funcDir := symbols.NewFunctionDirectory()
	constTable := symbols.NewConstantTable() // ← Ahora sin parámetros

	return &Parser{
		sourceCode:    sourceCode,
		debug:         debug,
		functionDir:   funcDir,
		constantTable: constTable,
	}
}

// Parse parses the source code and returns the symbol table and errors
func (p *Parser) Parse() (*symbols.FunctionDirectory, []string) {
	// Create input stream
	input := antlr.NewInputStream(p.sourceCode)

	// Create lexer
	lexer := grammar.NewBabyDuckLexer(input)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	parser := grammar.NewBabyDuckParser(stream)

	// Create directory builder
	p.directoryBuilder = NewDirectoryBuilder(p.debug, p.functionDir, p.constantTable)

	// Parse and walk the tree
	tree := parser.Program()

	// If debug is enabled, also use debug listener
	if p.debug {
		debugListener := NewDebugListener(true)
		antlr.ParseTreeWalkerDefault.Walk(debugListener, tree)
		fmt.Println("--- Iniciando análisis semántico ---")
	}

	// Walk with main listener
	antlr.ParseTreeWalkerDefault.Walk(p.directoryBuilder, tree)

	return p.functionDir, p.directoryBuilder.Errors
}

// GetDirectoryBuilder returns the directory builder for accessing quadruples
func (p *Parser) GetDirectoryBuilder() *DirectoryBuilder {
	return p.directoryBuilder
}
