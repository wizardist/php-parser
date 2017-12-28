package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BitwiseNot) Name() string {
	return "BitwiseNot"
}

type BitwiseNot struct {
	name string
	expr node.Node
}

func NewBitwiseNot(expression node.Node) node.Node {
	return BitwiseNot{
		"BitwiseNot",
		expression,
	}
}

func (n BitwiseNot) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
