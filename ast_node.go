package arel

type BaseNode struct{}

type AstNode interface {
	CreateTrue() TrueNode
	CreateFalse() FalseNode
	CreateTableAlias(*Table, string) TableAliasNode
	CreateStringJoin() StringJoinNode
	CreateInnerJoin() InnerJoinNode
	CreateOuterJoin() OuterJoinNode
	CreateAnd() AndNode
	CreateOn() OnNode
	CreateGrouping() GroupingNode
	CreateLower() NamedFunctionNode
}
