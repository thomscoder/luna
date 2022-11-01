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
	// Map instructions to their hex dump - empty string for all other nodes
	MapTo interface{}
}

type ExpressionNode struct {
	Type  string
	Value interface{}
}

type ExpressionParams struct {
	Type  string
	Value int
}
