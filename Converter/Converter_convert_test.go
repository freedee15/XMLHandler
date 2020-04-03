package Converter

import (
	"fmt"
	"github.com/cucumber/godog/gherkin"
	"github.com/freedee15/XMLHandler/Node"
)

func iCreateANodeTree() error {

	nodeTree = &Node.Tree{Label: "root"}

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

func theOutputShouldBe(arg1 *gherkin.DocString) error {

	if output != arg1.Content {
		return fmt.Errorf("expected:\n%s\ngot:\n%s", arg1.Content, output)
	}

	return nil

}
