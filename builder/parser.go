package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/symbols"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// Parser wraps the ANTLR parser and manages the compilation process
type Parser struct {
	// Input
	sourceCode string
	debug      bool

	// Core components
	directoryBuilder *DirectoryBuilder
	functionDir      *symbols.FunctionDirectory
	constantTable    *symbols.ConstantTable
}

// NewParser creates a new parser instance with initialized symbol tables
func NewParser(sourceCode string, debug bool) *Parser {
	return &Parser{
		sourceCode:    sourceCode,
		debug:         debug,
		functionDir:   symbols.NewFunctionDirectory(),
		constantTable: symbols.NewConstantTable(),
	}
}

// =============================================================================
// MAIN PARSING FUNCTIONALITY
// =============================================================================

// Parse processes the source code and returns the symbol table and any errors
func (p *Parser) Parse() (*symbols.FunctionDirectory, []string) {
	// Create ANTLR parser
	parser := p.createANTLRParser()

	// Create directory builder
	p.directoryBuilder = NewDirectoryBuilder(p.debug, p.functionDir, p.constantTable)

	// Parse the program
	tree := parser.Program()

	// Walk the parse tree
	p.walkParseTree(tree)

	return p.functionDir, p.directoryBuilder.Errors
}

// =============================================================================
// ANTLR SETUP
// =============================================================================

func (p *Parser) createANTLRParser() *grammar.BabyDuckParser {
	// Create input stream
	input := antlr.NewInputStream(p.sourceCode)

	// Create lexer
	lexer := grammar.NewBabyDuckLexer(input)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create and return parser
	return grammar.NewBabyDuckParser(stream)
}

// =============================================================================
// PARSE TREE WALKING
// =============================================================================

func (p *Parser) walkParseTree(tree grammar.IProgramContext) {
	if p.debug {
		p.walkWithDebugListener(tree)
	}

	p.walkWithMainListener(tree)
}

func (p *Parser) walkWithDebugListener(tree grammar.IProgramContext) {
	debugListener := NewDebugListener(true)
	antlr.ParseTreeWalkerDefault.Walk(debugListener, tree)
	fmt.Println("--- Starting semantic analysis ---")
}

func (p *Parser) walkWithMainListener(tree grammar.IProgramContext) {
	antlr.ParseTreeWalkerDefault.Walk(p.directoryBuilder, tree)
}

// =============================================================================
// PUBLIC INTERFACE
// =============================================================================

// GetDirectoryBuilder returns the directory builder for accessing quadruples
func (p *Parser) GetDirectoryBuilder() *DirectoryBuilder {
	return p.directoryBuilder
}

// GetFunctionDirectory returns the function directory
func (p *Parser) GetFunctionDirectory() *symbols.FunctionDirectory {
	return p.functionDir
}

// GetConstantTable returns the constant table
func (p *Parser) GetConstantTable() *symbols.ConstantTable {
	return p.constantTable
}

// GetQuadruples returns the generated quadruples
func (p *Parser) GetQuadruples() []interface{} {
	if p.directoryBuilder == nil || p.directoryBuilder.QuadVisitor == nil {
		return []interface{}{}
	}

	quads := p.directoryBuilder.QuadVisitor.GetQuadruples()
	result := make([]interface{}, len(quads))
	for i, quad := range quads {
		result[i] = quad
	}
	return result
}

// HasErrors returns true if there were compilation errors
func (p *Parser) HasErrors() bool {
	return p.directoryBuilder != nil && len(p.directoryBuilder.Errors) > 0
}

// GetErrors returns all compilation errors
func (p *Parser) GetErrors() []string {
	if p.directoryBuilder == nil {
		return []string{}
	}
	return p.directoryBuilder.Errors
}

// PrintResults prints compilation results including quadruples and constants
func (p *Parser) PrintResults() {
	if p.directoryBuilder != nil {
		if p.directoryBuilder.QuadVisitor != nil {
			p.directoryBuilder.QuadVisitor.PrintQuadruples()
		}
		if p.constantTable != nil {
			p.constantTable.Print()
		}
	}
}
