// Code generated from BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // BabyDuck

import "github.com/antlr4-go/antlr/v4"

// BabyDuckListener is a complete listener for a parse tree produced by BabyDuckParser.
type BabyDuckListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterVar_decl is called when entering the var_decl production.
	EnterVar_decl(c *Var_declContext)

	// EnterId_list is called when entering the id_list production.
	EnterId_list(c *Id_listContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterCycle is called when entering the cycle production.
	EnterCycle(c *CycleContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterElse_part is called when entering the else_part production.
	EnterElse_part(c *Else_partContext)

	// EnterPrint_stmt is called when entering the print_stmt production.
	EnterPrint_stmt(c *Print_stmtContext)

	// EnterPrintexpr is called when entering the printexpr production.
	EnterPrintexpr(c *PrintexprContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelational is called when entering the relational production.
	EnterRelational(c *RelationalContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterAddop is called when entering the addop production.
	EnterAddop(c *AddopContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMulop is called when entering the mulop production.
	EnterMulop(c *MulopContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterParexpr is called when entering the parexpr production.
	EnterParexpr(c *ParexprContext)

	// EnterFactorsign is called when entering the factorsign production.
	EnterFactorsign(c *FactorsignContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterFuncs is called when entering the funcs production.
	EnterFuncs(c *FuncsContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterParam_list is called when entering the param_list production.
	EnterParam_list(c *Param_listContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterFuncbody is called when entering the funcbody production.
	EnterFuncbody(c *FuncbodyContext)

	// EnterF_call is called when entering the f_call production.
	EnterF_call(c *F_callContext)

	// EnterArg_list is called when entering the arg_list production.
	EnterArg_list(c *Arg_listContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitVar_decl is called when exiting the var_decl production.
	ExitVar_decl(c *Var_declContext)

	// ExitId_list is called when exiting the id_list production.
	ExitId_list(c *Id_listContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitCycle is called when exiting the cycle production.
	ExitCycle(c *CycleContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitElse_part is called when exiting the else_part production.
	ExitElse_part(c *Else_partContext)

	// ExitPrint_stmt is called when exiting the print_stmt production.
	ExitPrint_stmt(c *Print_stmtContext)

	// ExitPrintexpr is called when exiting the printexpr production.
	ExitPrintexpr(c *PrintexprContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelational is called when exiting the relational production.
	ExitRelational(c *RelationalContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitAddop is called when exiting the addop production.
	ExitAddop(c *AddopContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMulop is called when exiting the mulop production.
	ExitMulop(c *MulopContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitParexpr is called when exiting the parexpr production.
	ExitParexpr(c *ParexprContext)

	// ExitFactorsign is called when exiting the factorsign production.
	ExitFactorsign(c *FactorsignContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitFuncs is called when exiting the funcs production.
	ExitFuncs(c *FuncsContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitParam_list is called when exiting the param_list production.
	ExitParam_list(c *Param_listContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitFuncbody is called when exiting the funcbody production.
	ExitFuncbody(c *FuncbodyContext)

	// ExitF_call is called when exiting the f_call production.
	ExitF_call(c *F_callContext)

	// ExitArg_list is called when exiting the arg_list production.
	ExitArg_list(c *Arg_listContext)
}
