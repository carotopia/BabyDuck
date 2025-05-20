// Code generated from BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // BabyDuck

import "github.com/antlr4-go/antlr/v4"

type BaseBabyDuckVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBabyDuckVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitVars(ctx *VarsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitVar_decl(ctx *Var_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitId_list(ctx *Id_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitCycle(ctx *CycleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitElse_part(ctx *Else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitPrint_stmt(ctx *Print_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitPrintexpr(ctx *PrintexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitRel_expr(ctx *Rel_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitRelop(ctx *RelopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitAdd_expr(ctx *Add_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitAddop(ctx *AddopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitMulop(ctx *MulopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitFuncs(ctx *FuncsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitParam_list(ctx *Param_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitFuncbody(ctx *FuncbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitF_call(ctx *F_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBabyDuckVisitor) VisitArg_list(ctx *Arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}
