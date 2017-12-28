package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n GreaterOrEqual) Name() string {
	return "GreaterOrEqual"
}

type GreaterOrEqual struct {
	BinaryOp
}

func NewGreaterOrEqual(variable node.Node, expression node.Node) node.Node {
	return GreaterOrEqual{
		BinaryOp{
			"BinaryGreaterOrEqual",
			variable,
			expression,
		},
	}
}

func (n GreaterOrEqual) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
