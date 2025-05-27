package builder

import (
	"BabyDuckCompiler/grammar"
	"fmt"
)

// DebugListener ayuda a identificar qué métodos se están llamando
type DebugListener struct {
	*grammar.BaseBabyDuckListener
	debug bool
}

func NewDebugListener(debug bool) *DebugListener {
	return &DebugListener{
		BaseBabyDuckListener: &grammar.BaseBabyDuckListener{},
		debug:                debug,
	}
}

func (d *DebugListener) log(message string) {
	if d.debug {
		fmt.Printf("[DEBUG] %s\n", message)
	}
}

// Implementa todos los métodos Enter/Exit para ver cuáles se llaman
func (d *DebugListener) EnterProgram(ctx *grammar.ProgramContext) {
	d.log("EnterProgram")
}

func (d *DebugListener) ExitProgram(ctx *grammar.ProgramContext) {
	d.log("ExitProgram")
}

func (d *DebugListener) EnterVar_decl(ctx *grammar.Var_declContext) {
	d.log(fmt.Sprintf("EnterVar_decl: %s", ctx.GetText()))
}

func (d *DebugListener) ExitVar_decl(ctx *grammar.Var_declContext) {
	d.log(fmt.Sprintf("ExitVar_decl: %s", ctx.GetText()))
}

func (d *DebugListener) EnterAssign(ctx *grammar.AssignContext) {
	d.log(fmt.Sprintf("EnterAssign: %s", ctx.GetText()))
}

func (d *DebugListener) ExitAssign(ctx *grammar.AssignContext) {
	d.log(fmt.Sprintf("ExitAssign: %s", ctx.GetText()))
}

func (d *DebugListener) EnterExpression(ctx *grammar.ExpressionContext) {
	d.log(fmt.Sprintf("EnterExpression: %s", ctx.GetText()))
}

func (d *DebugListener) ExitExpression(ctx *grammar.ExpressionContext) {
	d.log(fmt.Sprintf("ExitExpression: %s", ctx.GetText()))
}

func (d *DebugListener) EnterValue(ctx *grammar.ValueContext) {
	d.log(fmt.Sprintf("EnterValue: %s", ctx.GetText()))
}

func (d *DebugListener) ExitValue(ctx *grammar.ValueContext) {
	d.log(fmt.Sprintf("ExitValue: %s", ctx.GetText()))
}

func (d *DebugListener) EnterPrint_stmt(ctx *grammar.Print_stmtContext) {
	d.log(fmt.Sprintf("EnterPrint_stmt: %s", ctx.GetText()))
}

func (d *DebugListener) ExitPrint_stmt(ctx *grammar.Print_stmtContext) {
	d.log(fmt.Sprintf("ExitPrint_stmt: %s", ctx.GetText()))
}
