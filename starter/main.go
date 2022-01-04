package main

import (
	"fmt"
	"github.com/ACb0y/antlr_learn/starter/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type MyListener struct {
	parser.BaseArrayInitListener
}

func (m *MyListener) EnterInit(ctx *parser.InitContext)   {
	fmt.Printf("enterInit=%s\n", ctx.GetText())
}

func (m *MyListener) EnterValue(ctx *parser.ValueContext)   {
	fmt.Printf("enterValue=%s\n", ctx.INT().GetText())
 }

func main() {
	fmt.Println("test")
	// 创建一个charStream，文本作为输入
	input := antlr.NewInputStream("{1, 2, 3, 4, 5, 6}")
	// 创建一个词法分析器，处理输入的charStream
	lexer := parser.NewArrayInitLexer(input)
	// 使用lexer创建一个tokeStream
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// 创建一个语法分析器，处理tokens中的信息
	myParser := parser.NewArrayInitParser(tokens)
	// 针对init规则，开始语法分析，生成ast抽象语法树？
	tree := myParser.Init()

	walker := antlr.NewParseTreeWalker()
	myListener := MyListener{}
	walker.Walk(&myListener, tree)
}


