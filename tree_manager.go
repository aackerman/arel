package arel

type TreeManager struct {
	engine *Engine
	ctx    AstNode
	Ast    AstNode
}

func (t *TreeManager) ToSql() string {
	return t.Visitor().Accept(t.Ast)
}

func (t *TreeManager) Visitor() Visitor {
	return t.engine.Connection().Visitor()
}
