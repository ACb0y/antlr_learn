package main

import (
	"fmt"
	"github.com/ACb0y/antlr_learn/tour/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type MyVisitor struct {
	parser.BaseLabeledExprVisitor
	str2Int map[string]int32
}

func NewVisitor() *MyVisitor {
	return &MyVisitor{str2Int: make(map[string]int32)}
}

func (v *MyVisitor) Visit(tree antlr.ParseTree) int32 {
	switch val := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(val)
	case *parser.PrintExprContext:
		return v.VisitPrintExpr(val)
	case *parser.AssignContext:
		return v.VisitAssign(val)
	case *parser.BlankContext:
		return v.VisitBlank(val)
	case *parser.MulDivContext:
		return v.VisitMulDiv(val)
	case *parser.AddSubContext:
		return v.VisitAddSub(val)
	case *parser.IntContext:
		return v.VisitInt(val)
	case *parser.IdContext:
		return v.VisitId(val)
	case *parser.ParensContext:
		return v.VisitParens(val)
	default:
		panic("Unknown context")
	}
}

func (v *MyVisitor) VisitProg(ctx *parser.ProgContext) int32 {
	for _, value := range ctx.GetExpr_list() {
		v.Visit(value)
	}
	return 0
}

func (v *MyVisitor) VisitPrintExpr(ctx *parser.PrintExprContext) int32 {
	value := v.Visit(ctx.Expr())
	fmt.Println(value)
	return 0
}

func (v *MyVisitor) VisitAssign(ctx *parser.AssignContext) int32 {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.str2Int[id] = value
	return value
}

func (v *MyVisitor) VisitBlank(ctx *parser.BlankContext) int32 {
	// do nothing.
	return 0
}


func (v *MyVisitor) VisitParens(ctx *parser.ParensContext) int32 {
	return v.Visit(ctx.Expr())
}

func (v *MyVisitor) VisitMulDiv(ctx *parser.MulDivContext) int32 {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	if ctx.GetOp().GetTokenType() == parser.LabeledExprLexerMUL {
		return left * right
	} else {
		return left / right
	}
}

func (v *MyVisitor) VisitAddSub(ctx *parser.AddSubContext) int32 {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	if ctx.GetOp().GetTokenType() == parser.LabeledExprLexerADD {
		return left + right
	} else {
		return left - right
	}
}

func (v *MyVisitor) VisitId(ctx *parser.IdContext) int32 {
	id := ctx.ID().GetText()
	if value, ok := v.str2Int[id]; ok {
		return value
	}
	return 0
}

func (v *MyVisitor) VisitInt(ctx *parser.IntContext) int32 {
	value, _ := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	return int32(value)
}

func main() {
	fileName := "./input.txt"
	// 创建一个FileStream
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		fmt.Printf("open % failed. err=%s\n", fileName, err.Error())
		return
	}
	// 创建一个词法分析器，处理输入的charStream
	lexer := parser.NewLabeledExprLexer(input)
	// 使用lexer创建一个tokeStream
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// 创建一个语法分析器，处理tokens中的信息
	exprParser := parser.NewLabeledExprParser(tokens)
	exprParser.BuildParseTrees = true
	// 开始语法分析
	tree := exprParser.Prog()
	myVisitor := NewVisitor()
	// 遍历ast树
	myVisitor.Visit(tree)
}


