// Code generated from BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // BabyDuck

import "github.com/antlr4-go/antlr/v4"

// BaseBabyDuckListener is a complete listener for a parse tree produced by BabyDuckParser.
type BaseBabyDuckListener struct{}

var _ BabyDuckListener = &BaseBabyDuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBabyDuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBabyDuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBabyDuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBabyDuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBabyDuckListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBabyDuckListener) ExitProgram(ctx *ProgramContext) {}

// EnterVars is called when production vars is entered.
func (s *BaseBabyDuckListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BaseBabyDuckListener) ExitVars(ctx *VarsContext) {}

// EnterVar_decl is called when production var_decl is entered.
func (s *BaseBabyDuckListener) EnterVar_decl(ctx *Var_declContext) {}

// ExitVar_decl is called when production var_decl is exited.
func (s *BaseBabyDuckListener) ExitVar_decl(ctx *Var_declContext) {}

// EnterId_list is called when production id_list is entered.
func (s *BaseBabyDuckListener) EnterId_list(ctx *Id_listContext) {}

// ExitId_list is called when production id_list is exited.
func (s *BaseBabyDuckListener) ExitId_list(ctx *Id_listContext) {}

// EnterType is called when production type is entered.
func (s *BaseBabyDuckListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseBabyDuckListener) ExitType(ctx *TypeContext) {}

// EnterBody is called when production body is entered.
func (s *BaseBabyDuckListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseBabyDuckListener) ExitBody(ctx *BodyContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBabyDuckListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBabyDuckListener) ExitStatement(ctx *StatementContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseBabyDuckListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseBabyDuckListener) ExitAssign(ctx *AssignContext) {}

// EnterCycle is called when production cycle is entered.
func (s *BaseBabyDuckListener) EnterCycle(ctx *CycleContext) {}

// ExitCycle is called when production cycle is exited.
func (s *BaseBabyDuckListener) ExitCycle(ctx *CycleContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseBabyDuckListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseBabyDuckListener) ExitCondition(ctx *ConditionContext) {}

// EnterElse_part is called when production else_part is entered.
func (s *BaseBabyDuckListener) EnterElse_part(ctx *Else_partContext) {}

// ExitElse_part is called when production else_part is exited.
func (s *BaseBabyDuckListener) ExitElse_part(ctx *Else_partContext) {}

// EnterPrint_stmt is called when production print_stmt is entered.
func (s *BaseBabyDuckListener) EnterPrint_stmt(ctx *Print_stmtContext) {}

// ExitPrint_stmt is called when production print_stmt is exited.
func (s *BaseBabyDuckListener) ExitPrint_stmt(ctx *Print_stmtContext) {}

// EnterPrintexpr is called when production printexpr is entered.
func (s *BaseBabyDuckListener) EnterPrintexpr(ctx *PrintexprContext) {}

// ExitPrintexpr is called when production printexpr is exited.
func (s *BaseBabyDuckListener) ExitPrintexpr(ctx *PrintexprContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseBabyDuckListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseBabyDuckListener) ExitConstant(ctx *ConstantContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBabyDuckListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBabyDuckListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelational is called when production relational is entered.
func (s *BaseBabyDuckListener) EnterRelational(ctx *RelationalContext) {}

// ExitRelational is called when production relational is exited.
func (s *BaseBabyDuckListener) ExitRelational(ctx *RelationalContext) {}

// EnterRelop is called when production relop is entered.
func (s *BaseBabyDuckListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BaseBabyDuckListener) ExitRelop(ctx *RelopContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseBabyDuckListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseBabyDuckListener) ExitExp(ctx *ExpContext) {}

// EnterAddop is called when production addop is entered.
func (s *BaseBabyDuckListener) EnterAddop(ctx *AddopContext) {}

// ExitAddop is called when production addop is exited.
func (s *BaseBabyDuckListener) ExitAddop(ctx *AddopContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBabyDuckListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBabyDuckListener) ExitTerm(ctx *TermContext) {}

// EnterMulop is called when production mulop is entered.
func (s *BaseBabyDuckListener) EnterMulop(ctx *MulopContext) {}

// ExitMulop is called when production mulop is exited.
func (s *BaseBabyDuckListener) ExitMulop(ctx *MulopContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseBabyDuckListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseBabyDuckListener) ExitFactor(ctx *FactorContext) {}

// EnterParexpr is called when production parexpr is entered.
func (s *BaseBabyDuckListener) EnterParexpr(ctx *ParexprContext) {}

// ExitParexpr is called when production parexpr is exited.
func (s *BaseBabyDuckListener) ExitParexpr(ctx *ParexprContext) {}

// EnterFactorsign is called when production factorsign is entered.
func (s *BaseBabyDuckListener) EnterFactorsign(ctx *FactorsignContext) {}

// ExitFactorsign is called when production factorsign is exited.
func (s *BaseBabyDuckListener) ExitFactorsign(ctx *FactorsignContext) {}

// EnterValue is called when production value is entered.
func (s *BaseBabyDuckListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseBabyDuckListener) ExitValue(ctx *ValueContext) {}

// EnterFuncs is called when production funcs is entered.
func (s *BaseBabyDuckListener) EnterFuncs(ctx *FuncsContext) {}

// ExitFuncs is called when production funcs is exited.
func (s *BaseBabyDuckListener) ExitFuncs(ctx *FuncsContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseBabyDuckListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseBabyDuckListener) ExitFunc(ctx *FuncContext) {}

// EnterParam_list is called when production param_list is entered.
func (s *BaseBabyDuckListener) EnterParam_list(ctx *Param_listContext) {}

// ExitParam_list is called when production param_list is exited.
func (s *BaseBabyDuckListener) ExitParam_list(ctx *Param_listContext) {}

// EnterParam is called when production param is entered.
func (s *BaseBabyDuckListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseBabyDuckListener) ExitParam(ctx *ParamContext) {}

// EnterFuncbody is called when production funcbody is entered.
func (s *BaseBabyDuckListener) EnterFuncbody(ctx *FuncbodyContext) {}

// ExitFuncbody is called when production funcbody is exited.
func (s *BaseBabyDuckListener) ExitFuncbody(ctx *FuncbodyContext) {}

// EnterF_call is called when production f_call is entered.
func (s *BaseBabyDuckListener) EnterF_call(ctx *F_callContext) {}

// ExitF_call is called when production f_call is exited.
func (s *BaseBabyDuckListener) ExitF_call(ctx *F_callContext) {}

// EnterArg_list is called when production arg_list is entered.
func (s *BaseBabyDuckListener) EnterArg_list(ctx *Arg_listContext) {}

// ExitArg_list is called when production arg_list is exited.
func (s *BaseBabyDuckListener) ExitArg_list(ctx *Arg_listContext) {}
