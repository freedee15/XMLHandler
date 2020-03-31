package Node

import (
	"fmt"
	"strings"
)

type Tree struct {
	Label    string
	children []*Node
}

func (t *Tree) AddChild(child *Node) error {

	if t != child.parent {
		child.parent = t
		t.children = append(t.children, child)
	} else {
		return fmt.Errorf("\"%s\" is already parent of node \"%s\"", t.Label, child.label)
	}

	return nil

}

func (t *Tree) GetChildFromLabel(label string) (*Node, error) {

	if strings.Replace(label, " ", "", -1) == "" {
		return nil, fmt.Errorf("empty label")
	}

	for _, c := range t.children {

		if c.label == label {
			return c, nil
		}

	}

	return nil, fmt.Errorf("no child with label \"%s\"", label)

}
