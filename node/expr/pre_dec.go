package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n PreDec) Name() string {
	return "PreDec"
}

type PreDec struct {
	name     string
	variable node.Node
}

func NewPreDec(variableession node.Node) node.Node {
	return PreDec{
		"PreDec",
		variableession,
	}
}

func (n PreDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
