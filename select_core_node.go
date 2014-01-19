package rel

type SelectCoreNode struct {
	Source       *JoinSource
	Top          *TopNode
	Projections  *[]Visitable
	SetQuanifier Visitable
	Wheres       *[]Visitable
	Groups       *[]*GroupNode
	Having       *HavingNode
	Windows      *[]Visitable
	BaseVisitable
}

func NewSelectCoreNode() SelectCoreNode {
	return SelectCoreNode{
		Source: &JoinSource{
			Right: []Visitable{},
		},
	}
}

func (node *SelectCoreNode) SetFrom(v Visitable) {
	node.Source.Left = v
}
