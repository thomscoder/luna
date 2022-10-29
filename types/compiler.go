package types

type Matcher struct {
	Type  string
	Value string
}

type Token struct {
	Type  string
	Value string
	Index int
}

type AstNode struct {
	Type       string
	Expression ExpressionNode
}

type ExpressionNode struct {
	Type  string
	Value interface{}
}
