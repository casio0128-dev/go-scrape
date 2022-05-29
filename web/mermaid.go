package web

import (
	"fmt"
	"strings"
)

type (
	Mermaid struct {
		Nodes     []*Node
		Relations []*Relation
	}

	NodeType string
	Node     struct {
		Id      string
		Content string
		Type    NodeType
	}

	Relation struct {
		from         Node
		To           Node
		AssignSymbol Assign
	}

	AssignType string
	Assign     struct {
		Type    AssignType
		Comment string
	}
)

func TestNewMermaid() *Mermaid {
	var nodes []*Node
	var relations []*Relation

	nodes = append(nodes, NewNode("TEST1", "test1 node", Normal))
	nodes = append(nodes, NewNode("TEST2", "test2 node", Normal))
	nodes = append(nodes, NewNode("IF", "test3 node If", If))

	relations = append(relations, NewRelation(*nodes[0], *nodes[1], *NewAssign(NormalArrow1, "")))
	relations = append(relations, NewRelation(*nodes[1], *nodes[2], *NewAssign(NormalArrow1, "")))
	relations = append(relations, NewRelation(*nodes[2], *nodes[0], *NewAssign(NormalArrow1, "true")))

	return NewMermaid(nodes, relations)
}

func NewAssign(assignType AssignType, comment string) *Assign {
	return &Assign{Type: assignType, Comment: comment}
}

func NewRelation(from Node, to Node, assignSymbol Assign) *Relation {
	return &Relation{from: from, To: to, AssignSymbol: assignSymbol}
}

func NewNode(id string, content string, nodeType NodeType) *Node {
	return &Node{Id: id, Content: content, Type: nodeType}
}

func NewMermaid(nodes []*Node, relations []*Relation) *Mermaid {
	return &Mermaid{Nodes: nodes, Relations: relations}
}

// NodeType
const (
	Normal NodeType = "Normal"
	If     NodeType = "If"
)

// AssignType
const (
	Normal1 AssignType = "---"
	Normal2 AssignType = "----"
	Normal3 AssignType = "-----"

	NormalArrow1 AssignType = "-->"
	NormalArrow2 AssignType = "--->"
	NormalArrow3 AssignType = "---->"

	Thick1 AssignType = "==="
	Thick2 AssignType = "===="
	Thick3 AssignType = "====="

	ThickArrow1 AssignType = "==>"
	ThickArrow2 AssignType = "===>"
	ThickArrow3 AssignType = "====>"

	Dotted1 AssignType = "-.-"
	Dotted2 AssignType = "-..-"
	Dotted3 AssignType = "-...-"

	DottedArrow1 AssignType = "-.->"
	DottedArrow2 AssignType = "-..->"
	DottedArrow3 AssignType = "-...->"
)

func (m *Mermaid) String() string {
	var result string = "flowchart TB\n"
	for _, node := range m.Nodes {
		switch node.Type {
		case Normal:
			result += fmt.Sprintf("\t%s(%s)\n", node.Id, node.Content)
		case If:
			result += fmt.Sprintf("\t%s{%s}\n", node.Id, node.Content)
		default:
			panic(fmt.Errorf("Invalid node type %s", node.Type))
		}
	}
	result += "\n"

	for _, relation := range m.Relations {
		assign := relation.AssignSymbol
		symbol := string(assign.Type)
		if !strings.EqualFold(assign.Comment, "") {
			symbol = fmt.Sprintf("%s[%s]", symbol, assign.Comment)
		}
		result += "\t" + relation.from.Id + symbol + relation.To.Id + "\n"
	}
	return result
}
