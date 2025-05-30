package builder

import (
	"BabyDuckCompiler/grammar"
	"BabyDuckCompiler/symbols"
	"github.com/antlr4-go/antlr/v4"
)

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

// Parse - Método principal que maneja todo el proceso
func (p *PureVisitorParser) Parse() (*symbols.FunctionDirectory, []string) {
	// PASO 1: Crear parser ANTLR
	parser := p.createANTLRParser()

	// PASO 2: Primera pasada - Builder para tabla de símbolos
	builder := NewDirectoryBuilder(p.debug, p.functionDir, p.constantTable)
	tree := parser.Program()
	antlr.ParseTreeWalkerDefault.Walk(builder, tree)

	// PASO 3: Segunda pasada - Visitor puro para cuádruplos
	p.visitor = NewPureQuadrupleVisitor(p.functionDir, p.constantTable, p.debug)

	// CORRECCIÓN: Crear nuevo parser y hacer casting correcto
	parser2 := p.createANTLRParser()
	tree2 := parser2.Program()

	// Hacer casting seguro de IProgramContext a *grammar.ProgramContext
	if programCtx, ok := tree2.(*grammar.ProgramContext); ok && programCtx != nil {
		p.visitor.VisitProgram(programCtx)
	}

	// PASO 4: Combinar errores y mostrar resultados
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

// ========== INTERFAZ PÚBLICA ==========

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
