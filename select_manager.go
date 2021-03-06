package rel

import (
	"log"
)

type SelectManager struct {
	Engine Engine
	Ast    *SelectStatementNode
	Ctx    *SelectCoreNode
	BaseVisitable
}

func Select(visitables ...Visitable) *SelectManager {
	return NewSelectManager(RelEngine, &Table{}).Select(visitables...)
}

func NewSelectManager(engine Engine, table *Table) *SelectManager {
	if engine == nil {
		engine = RelEngine
	}
	stmt := NewSelectStatementNode()
	manager := SelectManager{
		Engine: engine,
		Ast:    stmt,
		Ctx:    stmt.Cores[len(stmt.Cores)-1],
	}
	// setup initial join source
	manager.From(table)
	return &manager
}

func (mgr *SelectManager) ToSql() string {
	return mgr.Engine.Visitor().Accept(mgr.Ast)
}

func (mgr *SelectManager) Project(visitables ...Visitable) *SelectManager {
	return mgr.Select(visitables...)
}

func (mgr *SelectManager) Select(visitables ...Visitable) *SelectManager {
	for _, selection := range visitables {
		if mgr.Ctx.Selections == nil {
			mgr.Ctx.Selections = &[]Visitable{}
		}

		*mgr.Ctx.Selections = append(*mgr.Ctx.Selections, selection)
	}
	return mgr
}

func (mgr *SelectManager) From(table interface{}) *SelectManager {
	var v Visitable
	switch t := table.(type) {
	case *Table:
		v = t
	case Table:
		v = &t
	case string:
		v = NewTable(t)
	}
	mgr.Ctx.Source.Left = v
	return mgr
}

func (mgr *SelectManager) As(name string) *TableAliasNode {
	return &TableAliasNode{
		Relation: &GroupingNode{Expr: []Visitable{mgr.Ast}},
		Name:     name,
	}
}

func (mgr *SelectManager) On(visitables ...Visitable) *SelectManager {
	right := mgr.Ctx.Source.Right

	if len(right) > 0 {
		last := right[len(right)-1]
		switch val := last.(type) {
		case *InnerJoinNode:
			val.Right = mgr.NewOnNode(mgr.collapse(visitables...))
		case *OuterJoinNode:
			val.Right = mgr.NewOnNode(mgr.collapse(visitables...))
		default:
			log.Fatalf("Unable to call On with input type %T", val)
		}
	}

	return mgr
}

func (mgr *SelectManager) Using(str string) *SelectManager {
	right := mgr.Ctx.Source.Right

	if len(right) > 0 {
		last := right[len(right)-1]
		switch val := last.(type) {
		case *InnerJoinNode:
			val.Right = &UsingNode{Expr: &QuotedNode{Raw: str}}
		case *OuterJoinNode:
			val.Right = &UsingNode{Expr: &QuotedNode{Raw: str}}
		default:
			log.Fatalf("Unable to call On with input type %T", val)
		}
	}

	return mgr
}

func (mgr *SelectManager) Join(visitable Visitable) *SelectManager {
	return mgr.InnerJoin(visitable)
}

func (mgr *SelectManager) InnerJoin(visitable Visitable) *SelectManager {
	mgr.Ctx.Source.Right = append(mgr.Ctx.Source.Right, &InnerJoinNode{Left: visitable})
	return mgr
}

func (mgr *SelectManager) OuterJoin(visitable Visitable) *SelectManager {
	mgr.Ctx.Source.Right = append(mgr.Ctx.Source.Right, &OuterJoinNode{Left: visitable})
	return mgr
}

func (mgr *SelectManager) Lock(node SqlLiteralNode) *SelectManager {
	mgr.Ast.Lock = NewLockNode(node)
	return mgr
}

func (mgr *SelectManager) LockForUpdate() *SelectManager {
	mgr.Ast.Lock = NewLockNode(Sql("FOR UPDATE"))
	return mgr
}

func (mgr *SelectManager) Take(i int) *SelectManager {
	return mgr.Limit(i)
}

func (mgr *SelectManager) Limit(i int) *SelectManager {
	mgr.Ast.Limit = NewLimitNode(Sql(i))
	return mgr
}

func (mgr *SelectManager) Exists() *ExistsNode {
	return NewExistsNode(mgr.Ast)
}

func (mgr *SelectManager) Order(visitables ...Visitable) *SelectManager {
	if len(visitables) > 0 {
		if mgr.Ast.Orders == nil {
			mgr.Ast.Orders = &[]Visitable{}
		}
		for _, v := range visitables {
			*mgr.Ast.Orders = append(*mgr.Ast.Orders, v)
		}
	}
	return mgr
}

func (mgr *SelectManager) Where(visitable Visitable) *SelectManager {
	if mgr.Ctx.Wheres == nil {
		mgr.Ctx.Wheres = &[]Visitable{}
	}

	if expr, ok := visitable.(SelectManager); ok {
		*mgr.Ctx.Wheres = append(*mgr.Ctx.Wheres, expr.Ast)
	} else {
		*mgr.Ctx.Wheres = append(*mgr.Ctx.Wheres, visitable)
	}

	return mgr
}

func (mgr *SelectManager) GroupBy(visitables ...Visitable) *SelectManager {
	return mgr.Group(visitables...)
}

func (mgr *SelectManager) Group(visitables ...Visitable) *SelectManager {
	if len(visitables) > 0 {
		if mgr.Ctx.Groups == nil {
			mgr.Ctx.Groups = &[]Visitable{}
		}
		for _, v := range visitables {
			*mgr.Ctx.Groups = append(*mgr.Ctx.Groups, NewGroupNode(v))
		}
	}
	return mgr
}

func (mgr *SelectManager) Intersect(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).Intersect(stmt1, stmt2)
}

func (mgr *SelectManager) Union(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).Union(stmt1, stmt2)
}

func (mgr *SelectManager) UnionAll(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).UnionAll(stmt1, stmt2)
}

func (mgr *SelectManager) Except(stmt1 Visitable, stmt2 Visitable) *MultiStatementManager {
	return NewMultiStatementManager(mgr.Engine).Except(stmt1, stmt2)
}

func (mgr *SelectManager) Skip(i int) *SelectManager {
	mgr.Ast.Offset = NewOffsetNode(Sql(i))
	return mgr
}

func (mgr *SelectManager) Offset(i int) *SelectManager {
	return mgr.Skip(i)
}

func (mgr *SelectManager) Having(visitables ...Visitable) *SelectManager {
	mgr.Ctx.Having = NewHavingNode(mgr.collapse(visitables...))
	return mgr
}

func (mgr *SelectManager) Distinct() *SelectManager {
	mgr.Ctx.SetQuantifier = &DistinctNode{}
	return mgr
}

func (mgr *SelectManager) NotDistinct() *SelectManager {
	mgr.Ctx.SetQuantifier = nil
	return mgr
}

func (mgr *SelectManager) With(visitable Visitable) *SelectManager {
	mgr.Ast.With = &WithNode{Expr: visitable}
	return mgr
}

func (mgr *SelectManager) WithRecursive(visitable Visitable) *SelectManager {
	mgr.Ast.With = &WithRecursiveNode{Expr: visitable}
	return mgr
}

func (mgr *SelectManager) Window(node SqlLiteralNode) *NamedWindowNode {
	if mgr.Ctx.Windows == nil {
		mgr.Ctx.Windows = &[]Visitable{}
	}
	window := &NamedWindowNode{Name: node}
	*mgr.Ctx.Windows = append(*mgr.Ctx.Windows, window)
	return window
}

func (mgr *SelectManager) collapse(visitables ...Visitable) Visitable {
	var v Visitable

	// use the first Node if there is only one
	// else create and And node
	if len(visitables) == 1 {
		v = visitables[0]
	} else {
		v = mgr.NewAndNode(visitables...)
	}
	return v
}
