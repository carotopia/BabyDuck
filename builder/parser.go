package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/symbols"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// 🔧 CustomErrorCollector - detecta errores de sintaxis
type CustomErrorCollector struct {
	*antlr.DefaultErrorListener
	HasErrors bool
	Errors    []string
}

func NewCustomErrorCollector() *CustomErrorCollector {
	return &CustomErrorCollector{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		HasErrors:            false,
		Errors:               make([]string, 0),
	}
}

func (c *CustomErrorCollector) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// Llamar al default para que siga mostrando los errores como antes
	c.DefaultErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)

	// Marcar que hubo errores
	c.HasErrors = true
	errorMsg := fmt.Sprintf("Syntax error at line %d:%d - %s", line, column, msg)
	c.Errors = append(c.Errors, errorMsg)
}

// PureVisitorParser parser para el nuevo sistema de cuadruplos
type PureVisitorParser struct {
	sourceCode    string
	debug         bool
	visitor       *PureQuadrupleVisitor
	functionDir   *symbols.FunctionDirectory
	constantTable *symbols.ConstantTable
}

// NewPureVisitorParser crea el parser con la nueva arquitectura
func NewPureVisitorParser(sourceCode string, debug bool) *PureVisitorParser {
	return &PureVisitorParser{
		sourceCode:    sourceCode,
		debug:         debug,
		functionDir:   symbols.NewFunctionDirectory(),
		constantTable: symbols.NewConstantTable(),
	}
}

//  Parse - Método principal que maneja todo el proceso CON MANEJO DE ERRORES
func (p *PureVisitorParser) Parse() (*symbols.FunctionDirectory, []string) {
	// PASO 1: Crear parser ANTLR CON ERROR COLLECTOR
	input := antlr.NewInputStream(p.sourceCode)
	lexer := grammar.NewBabyDuckLexer(input)


	errorCollector := NewCustomErrorCollector()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorCollector)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewBabyDuckParser(stream)


	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorCollector)


	builder := NewDirectoryBuilder(p.debug, p.functionDir, p.constantTable)
	tree := parser.Program()


	if errorCollector.HasErrors {
		return nil, errorCollector.Errors
	}

	antlr.ParseTreeWalkerDefault.Walk(builder, tree)


	if len(builder.Errors) > 0 {
		return p.functionDir, builder.Errors
	}

	p.visitor = NewPureQuadrupleVisitor(p.functionDir, p.constantTable, p.debug)

	parser2 := p.createANTLRParser()
	tree2 := parser2.Program()

	if programCtx, ok := tree2.(*grammar.ProgramContext); ok && programCtx != nil {
		p.visitor.VisitProgram(programCtx)
	}

	allErrors := append(builder.Errors, p.visitor.GetErrors()...)

	if p.debug {
		p.visitor.PrintQuadruples()
		p.constantTable.Print()
	}

	return p.functionDir, allErrors
}

func (p *PureVisitorParser) createANTLRParser() *grammar.BabyDuckParser {
	input := antlr.NewInputStream(p.sourceCode)
	lexer := grammar.NewBabyDuckLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return grammar.NewBabyDuckParser(stream)
}



func (p *PureVisitorParser) GetFunctionDirectory() *symbols.FunctionDirectory {
	return p.functionDir
}

func (p *PureVisitorParser) GetConstantTable() *symbols.ConstantTable {
	return p.constantTable
}

func (p *PureVisitorParser) GetQuadruples() []interface{} {
	if p.visitor == nil {
		return []interface{}{}
	}

	quads := p.visitor.GetQuadruples()
	result := make([]interface{}, len(quads))
	for i, quad := range quads {
		result[i] = quad
	}
	return result
}

func (p *PureVisitorParser) HasErrors() bool {
	return p.visitor != nil && p.visitor.HasErrors()
}

func (p *PureVisitorParser) GetErrors() []string {
	if p.visitor == nil {
		return []string{}
	}
	return p.visitor.GetErrors()
}

func (p *PureVisitorParser) PrintResults() {
	if p.visitor != nil {
		p.visitor.PrintQuadruples()
		if p.constantTable != nil {
			p.constantTable.Print()
		}
	}
}
