package Node

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	children []*Node
	data     string
	label    string
	parent   Parent
}

func (n *Node) AddChild(child *Node) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	if n != child.GetParent() {
		child.SetParent(n)
		n.children = append(n.children, child)
	}
	return nil

}

func (n *Node) GetChildFromLabel(label string) (*Node, error) {

	if strings.Replace(label, " ", "", -1) == "" {
		return nil, fmt.Errorf("empty label")
	}

	for _, c := range n.children {

		if c.label == label {
			return c, nil
		}

	}

	return nil, fmt.Errorf("no child with label \"%s\"", label)

}

func (n *Node) GetData() string {

	return n.data

}

func (n *Node) GetDataAsBool() (bool, error) {

	if b, e := strconv.ParseBool(n.data); e != nil {
		return false, e
	} else {
		return b, nil
	}

}

func (n *Node) GetDataAsFloat() (float64, error) {

	if f, e := strconv.ParseFloat(n.data, 64); e != nil {
		return 0, e
	} else {
		return f, nil
	}

}

func (n *Node) GetDataAsInt() (int, error) {

	if i, e := strconv.Atoi(n.data); e != nil {
		return 0, e
	} else {
		return i, nil
	}

}

func (n *Node) GetLabel() string {

	return n.label

}

func (n *Node) GetParent() Parent {

	return n.parent

}

func (n *Node) SetData(s string) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = s
	return nil

}

func (n *Node) SetDataFromBool(b bool) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = strconv.FormatBool(b)
	return nil

}

func (n *Node) SetDataFromFloat(f float64) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = strconv.FormatFloat(f, 'f', -1, 64)
	return nil

}

func (n *Node) SetDataFromInt(i int) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = strconv.Itoa(i)
	return nil

}

func (n *Node) SetLabel(s string) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.label = s
	return nil

}

func (n *Node) SetParent(p Parent) {

	n.parent = p
	p.AddChild(n)

}
