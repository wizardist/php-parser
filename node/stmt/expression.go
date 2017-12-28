package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Expression) Name() string {
	return "Expression"
}

type Expression struct {
	name string
	expr node.Node
}

func NewExpression(expr node.Node) node.Node {
	return Expression{
		"Expression",
		expr,
	}
}

func (n Expression) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
