// Code generated from BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // BabyDuck

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BabyDuckParser.
type BabyDuckVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BabyDuckParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#vars.
	VisitVars(ctx *VarsContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#var_decl.
	VisitVar_decl(ctx *Var_declContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#id_list.
	VisitId_list(ctx *Id_listContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#cycle.
	VisitCycle(ctx *CycleContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#else_part.
	VisitElse_part(ctx *Else_partContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#print_stmt.
	VisitPrint_stmt(ctx *Print_stmtContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#printexpr.
	VisitPrintexpr(ctx *PrintexprContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#rel_expr.
	VisitRel_expr(ctx *Rel_exprContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#relop.
	VisitRelop(ctx *RelopContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#add_expr.
	VisitAdd_expr(ctx *Add_exprContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#addop.
	VisitAddop(ctx *AddopContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#mulop.
	VisitMulop(ctx *MulopContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#funcs.
	VisitFuncs(ctx *FuncsContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#param_list.
	VisitParam_list(ctx *Param_listContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#funcbody.
	VisitFuncbody(ctx *FuncbodyContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#f_call.
	VisitF_call(ctx *F_callContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#arg_list.
	VisitArg_list(ctx *Arg_listContext) interface{}
}
