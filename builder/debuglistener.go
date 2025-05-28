package builder

import (
	"BabyDuckCompiler/grammar"
	"fmt"
)

// DebugListener helps identify which methods are being called during parsing
type DebugListener struct {
	*grammar.BaseBabyDuckListener
	enabled bool
}

// NewDebugListener creates a new debug listener
func NewDebugListener(enabled bool) *DebugListener {
	return &DebugListener{
		BaseBabyDuckListener: &grammar.BaseBabyDuckListener{},
		enabled:              enabled,
	}
}

// =============================================================================
// CORE LOGGING FUNCTIONALITY
// =============================================================================

func (d *DebugListener) log(message string) {
	if d.enabled {
		fmt.Printf("[DEBUG] %s\n", message)
	}
}

func (d *DebugListener) logWithContext(event, context string) {
	if d.enabled {
		fmt.Printf("[DEBUG] %s: %s\n", event, context)
	}
}

// =============================================================================
// PROGRAM STRUCTURE DEBUGGING
// =============================================================================

func (d *DebugListener) EnterProgram(ctx *grammar.ProgramContext) {
	d.log("EnterProgram")
}

func (d *DebugListener) ExitProgram(ctx *grammar.ProgramContext) {
	d.log("ExitProgram")
}

// =============================================================================
// VARIABLE DECLARATION DEBUGGING
// =============================================================================

func (d *DebugListener) EnterVar_decl(ctx *grammar.Var_declContext) {
	d.logWithContext("EnterVar_decl", ctx.GetText())
}

func (d *DebugListener) ExitVar_decl(ctx *grammar.Var_declContext) {
	d.logWithContext("ExitVar_decl", ctx.GetText())
}

// =============================================================================
// ASSIGNMENT DEBUGGING
// =============================================================================

func (d *DebugListener) EnterAssign(ctx *grammar.AssignContext) {
	d.logWithContext("EnterAssign", ctx.GetText())
}

func (d *DebugListener) ExitAssign(ctx *grammar.AssignContext) {
	d.logWithContext("ExitAssign", ctx.GetText())
}

// =============================================================================
// EXPRESSION DEBUGGING
// =============================================================================

func (d *DebugListener) EnterExpression(ctx *grammar.ExpressionContext) {
	d.logWithContext("EnterExpression", ctx.GetText())
}

func (d *DebugListener) ExitExpression(ctx *grammar.ExpressionContext) {
	d.logWithContext("ExitExpression", ctx.GetText())
}

func (d *DebugListener) EnterValue(ctx *grammar.ValueContext) {
	d.logWithContext("EnterValue", ctx.GetText())
}

func (d *DebugListener) ExitValue(ctx *grammar.ValueContext) {
	d.logWithContext("ExitValue", ctx.GetText())
}

// =============================================================================
// STATEMENT DEBUGGING
// =============================================================================

func (d *DebugListener) EnterPrint_stmt(ctx *grammar.Print_stmtContext) {
	d.logWithContext("EnterPrint_stmt", ctx.GetText())
}

func (d *DebugListener) ExitPrint_stmt(ctx *grammar.Print_stmtContext) {
	d.logWithContext("ExitPrint_stmt", ctx.GetText())
}

// =============================================================================
// UTILITY METHODS
// =============================================================================

// SetEnabled allows toggling debug output during runtime
func (d *DebugListener) SetEnabled(enabled bool) {
	d.enabled = enabled
}

// IsEnabled returns current debug state
func (d *DebugListener) IsEnabled() bool {
	return d.enabled
}
