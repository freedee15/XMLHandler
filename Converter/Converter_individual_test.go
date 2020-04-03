package Converter

import (
	"XMLHandler/Node"
	"fmt"
)

func iCreateANodeTree() error {

	nodeTree = &Node.Tree{Label: "root"}

	return nil

}

func iCreateAChildNodeOfNodeTreeLabelled(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	n := &Node.Node{}

	if e := nodeTree.AddChild(n); e != nil {
		return e
	}

	if e := n.SetLabel(arg1); e != nil {
		return e
	}

	return nil

}

func iCreateAChildNodeOfNodeLabelled(arg1, arg2 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	c := &Node.Node{}

	n, e := nodeTree.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	if e = n.AddChild(c); e != nil {
		return e
	}
	if e = c.SetLabel(arg2); e != nil {
		return e
	}

	return nil

}

func iSelectChildNodeOfNodeTree(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	s, e := nodeTree.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	selected = s

	return nil

}

func iSelectChildNodeOfSelectedNode(arg1 string) error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	s, e := selected.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	selected = s

	return nil

}

func iConvertTheSelectedNodeToXML() error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	o, e := ConvertNodeToXML(selected)
	if e != nil {
		return e
	}
	output = o

	return nil

}

func theOutputShouldBe(arg1 string) error {

	if output != arg1 {
		return fmt.Errorf("expected %s, got %s", arg1, output)
	}

	return nil

}

func iSetTheSelectedNodesDataTo(arg1 string) error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	if e := selected.SetData(arg1); e != nil {
		return e
	}

	return nil

}
