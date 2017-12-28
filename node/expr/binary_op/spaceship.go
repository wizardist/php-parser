package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Spaceship) Name() string {
	return "Spaceship"
}

type Spaceship struct {
	BinaryOp
}

func NewSpaceship(variable node.Node, expression node.Node) node.Node {
	return Spaceship{
		BinaryOp{
			"BinarySpaceship",
			variable,
			expression,
		},
	}
}

func (n Spaceship) Walk(v node.Visitor) {
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
